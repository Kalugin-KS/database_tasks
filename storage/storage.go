package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

// NewStorage принимает на вход строку подключения к БД и возвращает новый пул подключений к БД
func NewStorage(conn string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	s := Storage{
		db: db,
	}
	return &s, nil

}

// Сруктура с основными параметрами задачи
type Task struct {
	ID       int
	Opened   int
	AuthorID int
	Title    string
	Content  string
}

// TasksByAuthor возвращает список задач из БД по автору
func (s *Storage) TasksByAuthor(authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT id, opened, author_id, title, content FROM tasks
	WHERE author_id = $1 ORDER BY id;`, authorID)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.AuthorID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

// NewTask создает новую задачу и возращает ее номер id
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(),
		`INSERT INTO tasks(author_id, title, content) VALUES
	($1, $2, $3) RETURNING id;`, t.AuthorID, t.Title, t.Content).Scan(&id)
	return id, err
}

// Tasks возвращает список всех задач
func (s *Storage) Tasks() ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `SELECT * FROM tasks ORDER BY id;`)
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.AuthorID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

// TasksByLabel возвращает список всех задач по метке
func (s *Storage) TasksByLabel(label string) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `SELECT 
	tasks.id, tasks.opened, tasks.author_id, tasks.title, tasks.content 
	FROM tasks, labels, tasks_labels WHERE 
	tasks.id = task_id AND labels.id = label_id AND labels.name = $1 ORDER BY id`, label)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.AuthorID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

// UpdateTasks обновляет задачу по id
func (s *Storage) UpdateTask(id int, title, content string) error {
	_, err := s.db.Exec(context.Background(), `UPDATE tasks 
	SET title = $1, content = $2 WHERE id = $3;`, title, content, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask удаляет задачу по id
func (s *Storage) DeleteTask(id int) error {
	// объединим две операции с БД в одну используя транзакцию
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), `DELETE FROM tasks_labels WHERE task_id = $1`, id)
	if err != nil {
		// откат транзакции в случае ошибки
		tx.Rollback(context.Background())
		return err
	}
	_, err = tx.Exec(context.Background(), `DELETE FROM tasks WHERE tasks.id = $1`, id)
	if err != nil {
		// откат транзакции в случае ошибки
		tx.Rollback(context.Background())
		return err
	}
	// фиксация (подтверждение) транзакции
	tx.Commit(context.Background())
	return nil
}
