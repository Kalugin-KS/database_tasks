DROP TABLE IF EXISTS tasks_labels, labels, tasks, users;

--таблица пользователей
CREATE TABLE users(
	id SERIAL PRIMARY KEY, 
	name TEXT NOT NULL
);

--таблица меток
CREATE TABLE labels(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

--таблица задач
CREATE TABLE tasks(
	id SERIAL PRIMARY KEY,
	opened BIGINT NOT NULL DEFAULT extract(epoch from now()),
	author_id INT REFERENCES users(id),
	title TEXT,
	content TEXT
);

--таблица связи между метками и задачами
CREATE TABLE tasks_labels(
	task_id INT REFERENCES tasks(id),
	label_id INT REFERENCES labels(id)
);

--заполнение таблиц исходными данными
INSERT INTO users(name) VALUES
	('Pasha'), 
	('Nick'),
	('Sam');
	
INSERT INTO labels(name) VALUES
	('mongo'), 
	('postgres'),
	('sql');

INSERT INTO tasks(author_id, title, content) VALUES
	(1, 'postgres', 'create DB postgreSQL'),
	(2, 'mongo', 'create DB mongo'),
	(1, 'sql', 'create DB mySQL');
	
INSERT INTO tasks_labels(task_id, label_id) VALUES
	(1, 2),
	(2, 1),
	(3, 3);

SELECT * FROM tasks;