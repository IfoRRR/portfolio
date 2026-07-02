# Портфолио — Зайцев Артем

Сайт-портфолио на **Vue 3 + Go** с REST API, автозагрузкой проектов с GitHub и аналитикой кликов.

## Структура

```
portfolio/
├── backend/          ← Go REST API
│   ├── main.go
│   ├── config.json   ← твои данные (имя, навыки, контакты)
│   └── assets/resume.pdf
└── frontend/         ← Vue 3 + Vite
    └── src/
```

## Быстрый старт

### 1. Бэкенд

```powershell
cd backend
go run .
```

API: http://localhost:8080

### 2. Фронтенд (режим разработки)

```powershell
cd frontend
npm install
npm run dev
```

Сайт: http://localhost:5173 (запросы к API проксируются на :8080)

### 3. Production (один сервер)

```powershell
cd frontend
npm run build

cd ../backend
go run .
```

Go отдаёт и API, и собранный Vue из `frontend/dist`.

## REST API

| Метод | URL | Описание |
|-------|-----|----------|
| GET | `/api/about` | Имя, описание, навыки, контакты |
| GET | `/api/projects` | Проекты с GitHub |
| GET | `/api/skills` | Навыки |
| GET | `/api/contact` | Контакты |
| GET | `/api/analytics` | Статистика кликов и скачиваний |
| POST | `/api/analytics/click` | Записать клик `{ "link": "...", "type": "project" }` |
| GET | `/api/resume` | Скачать резюме (+ счётчик) |

## Аналитика

Каждый клик по проекту, контакту или скачивание резюме сохраняется в `backend/analytics.json`.
На сайте есть раздел **«Статистика»** с live-данными.

## Мои проекты

- [WEB-socket](https://github.com/IfoRRR/WEB-socket) — WebSocket чат на Go + SQLite
- [WEB-Socket-2.0](https://github.com/IfoRRR/WEB-Socket-2.0) — Vue фронтенд для чата
- [URL-SHORTENER](https://github.com/IfoRRR/URL-SHORTENER) — Fullstack Go + Vue
