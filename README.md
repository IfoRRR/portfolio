# Портфолио — Зайцев Артем

Сайт-портфолио backend/frontend-разработчика на **Vue 3 + Go** с REST API, автоматической загрузкой проектов с GitHub и аналитикой кликов.

## Избранные проекты

| Проект | Что показывает |
|---|---|
| [vk-feed](https://github.com/IfoRRR/vk-feed) | Go, PostgreSQL, Redis, cursor pagination, идемпотентные лайки |
| [vk-music-demo](https://github.com/IfoRRR/vk-music-demo) | worker pool, Redis, PostgreSQL, REST API, graceful shutdown |
| [grpc-auth](https://github.com/IfoRRR/grpc-auth) | JWT, REST, gRPC interceptor и metadata |
| [WEB-socket](https://github.com/IfoRRR/WEB-socket) | WebSocket, комнаты, SQLite, concurrent Go |
| [photo-editor](https://github.com/IfoRRR/photo-editor) | Vue, Canvas и client-side обработка изображений |
| [URL-SHORTENER](https://github.com/IfoRRR/URL-SHORTENER) | Fullstack Go + Vue, короткие ссылки и аналитика |

## Структура

```text
portfolio/
├── backend/          ← Go REST API
│   ├── main.go
│   ├── config.json   ← данные профиля и контакты
│   └── assets/resume.pdf
└── frontend/         ← Vue 3 + Vite
    └── src/
```

## Быстрый старт

### Backend

```bash
cd backend
go run .
```

### Frontend в development

```bash
cd frontend
npm install
npm run dev
```

Сайт: `http://localhost:5173`  
API: `http://localhost:8080`

### Production

```bash
cd frontend
npm install
npm run build
cd ../backend
go run .
```

Go отдаёт REST API и собранный Vue frontend.

## REST API

| Метод | URL | Описание |
|---|---|---|
| `GET` | `/api/about` | профиль, описание и контакты |
| `GET` | `/api/projects` | список проектов с GitHub |
| `GET` | `/api/skills` | навыки |
| `GET` | `/api/contact` | контакты |
| `GET` | `/api/analytics` | статистика просмотров и кликов |
| `POST` | `/api/analytics/click` | запись клика |
| `GET` | `/api/resume` | скачивание резюме |

## Аналитика

Клики по проектам, контактам и скачивания резюме сохраняются локально в backend. Секреты и `.env` не должны попадать в публичный репозиторий.

## Направление

Основной фокус портфолио: **Go backend, PostgreSQL, Redis, Docker, gRPC, WebSocket и Vue**.
