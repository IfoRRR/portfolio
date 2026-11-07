# Портфолио — Зайцев Артем

Сайт-портфолио backend/frontend-разработчика на **Vue 3 + Go** с REST API, автоматической загрузкой проектов с GitHub и аналитикой кликов.

## Избранные проекты

| Проект | Что показывает |
|---|---|
| [WEB-socket](https://github.com/IfoRRR/WEB-socket) | Полный WebSocket-чат: Go, Vue, комнаты и SQLite |
| [URL-SHORTENER](https://github.com/IfoRRR/URL-SHORTENER) | Fullstack Go + Vue, короткие ссылки и аналитика |
| [vk-quiz](https://github.com/IfoRRR/vk-quiz) | WebSocket-квиз с комнатами, таймером и табло |
| [vk-kanban](https://github.com/IfoRRR/vk-kanban) | Командная Kanban-доска с ролями и realtime-обновлениями |
| [meeting-room-mvp](https://github.com/IfoRRR/meeting-room-mvp) | MVP бронирования переговорных комнат |

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

### Docker production

Для единого входа через nginx нужны соседние каталоги `portfolio`,
`websocket`, `url-shortener`, `vk-quiz`, `vk-kanban` и `meeting-room-mvp`:

```bash
cd portfolio/deploy
docker compose build
docker compose up -d
```

Сайт будет доступен через `http://IP:8081/`, переговорная — через
`http://IP:8081/meeting/`. Наружу публикуется только nginx; порты приложений и
PostgreSQL остаются внутри Docker-сети. Конфигурация деплоя находится в
[`deploy/`](deploy/).

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
