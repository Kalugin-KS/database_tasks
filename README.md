# database_tasks
Данный пакет предоставляет необходимые методы для работы с БД tasks. 

********************
БД tasks содержит следующие таблицы:

- Пользователи
- Метки
- Задачи
- Таблица связей между задачами и метками
********************
API пакета storage позволяет:

- Создавать новые задачи
- Получать список всех задач
- Получать список задач по автору
- Получать список задач по метке
- Обновлять задачу по id
- Удалять задачу по id
*********************
Схема базы данных с исходными данными - `scheme.sql`
*********************
Запуск программы:  Сначала создаем базу данных postgresql и выполняем запрос как в файле `scheme.sql`, затем выполняем команду в терминале dbpass=*your_password* go run main.go

*your_password* - ваш пароль подключения к БД
