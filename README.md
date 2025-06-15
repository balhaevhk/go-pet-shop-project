# 🐾 Go Pet Shop

Интернет-магазин товаров для питомцев на Go с использованием PostgreSQL.

Проект создан для отработки навыков:

- Проектирования REST API
- Работы с SQL (SELECT, INSERT, JOIN, TRANSACTION)
- Организации backend-архитектуры (handlers, storage, config)
- Управления миграциями
- Написания многослойного Go-приложения

---

## 🧩 Стек технологий

- Go 1.22+
- PostgreSQL
- `net/http`, `database/sql`, `github.com/lib/pq`
- `golang-migrate` — миграции
- `golangci-lint` — линтинг
- `Taskfile` — автоматизация запуска

---

## 🚀 Как запустить


```bash
    1. Клонируй репозиторий
git clone https://github.com/your-username/go-pet-shop.git

cd go-pet-shop

    2. Создай .env файл

DATABASE_URL=postgres://user:password@localhost:5432/petshop?sslmode=disable

    3. Запусти миграции

task migrate

    4. Запусти сервер

go run cmd/app/main.go

    5. Проверь, что всё работает

curl http://localhost:8080/status
# Response: {"status": "ok"}

⸻

🧪 Проверка кода

task linter

⸻

🧱 Структура проекта

go-pet-shop/
├── cmd/                # Точки входа: app и migrator
├── config/             # YAML-конфиги
├── internal/           # Внутренняя логика
│   ├── config/         # Чтение/валидация конфига
│   ├── handlers/       # HTTP-обработчики
│   ├── storage/        # Работа с БД (PostgreSQL)
│   └── lib/            # Вспомогательные пакеты
├── migrations/         # SQL-миграции
├── models/             # DTO и доменные модели
├── storage             # Реализация работы с БД
├── .env                # Переменные окружения
├── Taskfile.yaml       # Task runner
└── README.md           # Ты сейчас тут

⸻

🧾 Описание версий

✅ Версия 1 — Пользователи и Товары

🔹 Функциональность:
	•	CRUD-операции с таблицами users и products
	•	Простые SELECT и INSERT
	•	Начальная структура проекта

🔹 Роуты:
	•	POST /users
	•	GET /users
	•	GET /users/email
	•	POST /products
	•	GET /products
	•	GET /products/{id}
	•	PUT /products/{id}
	•	DELETE /products/{id}

🔹 Ветка: v1

⸻

✅ Версия 2 — Заказы и позиции заказа

🔹 Функциональность:
	•	Добавление заказов и позиций
	•	JOIN-запросы для заказов

🔹 Роуты:
	•	POST /orders
	•	POST /orders/{id}/items
	•	GET /orders/{id}
	•	GET /users/orders

🔹 Ветка: v2

⸻

✅ Версия 3 — Оформление заказа (транзакции)

🔹 Функциональность:
	•	Транзакционное оформление заказа (PlaceOrder)
	•	Обновление stock
	•	Создание записи transactions

🔹 Роуты:
	•	POST /checkout

🔹 Ветка: v3

⸻

✅ Версия 4 — История и аналитика

🔹 Функциональность:
	•	История заказов пользователя (JOIN + вложенные SELECT)
	•	Аналитика: популярные товары (GROUP BY + SUM)

🔹 Роуты:
	•	GET /users/history
	•	GET /products/popular

🔹 Ветка: v4

⸻

🛠️ Используемые команды (Taskfile)

    Команда	                    Описание
task migrate	    Запуск миграций через cmd/migrator
task linter         Линтинг кода через golangci-lint