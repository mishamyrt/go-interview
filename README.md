# Задание для собеседования Go–разработчика

Данный репозиторий содержит тестовое задание для собеседования в формате код–ревью.

## Дано

Разработчику дали задачу написать бэкенд для todo-сервиса (управление списком задач). Сервис должен уметь:

* Создавать задачи
* Редактировать задачи
* Удалять задачи
* Получать список задач
* Получать задачу по ID
* Экспортировать задачи в файл

Разработчик реализовал веб–сервер, покрыл тестами и в один коммит отправил его на ревью.

## Задача

Найти проблемные места, объяснить что именно с ними не так и предложить решение. Время выполнениям ~70 минут.

## Рекомендуемый порядок просмотра файлов

1. [main.go](main.go)
2. [handler.go](handler.go)
3. [db.go](db.go)
4. [utils.go](utils.go)
5. [utils_test.go](utils_test.go)

## Чек–лист

Если вы ревьюер и вам пригодится готовый чек–лист по этому заданию — пишите в [телеграм](https://t.me/mishamyrt).