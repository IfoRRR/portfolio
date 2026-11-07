# Production deployment

The compose file expects these repositories as sibling directories:

```text
/opt/
├── portfolio/
├── vk-quiz/
├── vk-kanban/
├── websocket/
├── url-shortener/
└── meeting-room-mvp/
```

Run from this directory:

```bash
cp .env.example .env
# replace the example database password before starting
docker compose build
docker compose up -d
```

Only nginx publishes `8081:80`. Application ports remain inside the Docker network.

The current compose stack deploys the portfolio, WebSocket chat, shortener,
quiz, Kanban board and meeting-room MVP. The meeting room is available at
`/meeting/` and uses an internal PostgreSQL service; neither PostgreSQL nor the
application port is published to the VPS.

For the first deployment, copy the resume into the `portfolio-resume` volume
or use the bundled `backend/assets/resume.pdf` as the initial seed. Keep real
`.env` files and application data outside Git.
