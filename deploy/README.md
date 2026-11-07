# Production deployment

The compose file expects these repositories as sibling directories:

```text
/opt/
├── portfolio/
├── vk-quiz/
├── vk-kanban/
├── websocket/
└── url-shortener/
└── meeting-room-mvp/  # optional until its private source is available on the server
```

Run from this directory:

```bash
docker compose build
docker compose up -d
```

Only nginx publishes `8081:80`. Application ports remain inside the Docker network.

The current compose stack deploys the portfolio, WebSocket chat, shortener,
quiz and Kanban board. The meeting-room project is shown as a portfolio case,
but is not exposed as a live service yet because its repository is private and
requires authenticated checkout on the VPS.

For the first deployment, copy the resume into the `portfolio-resume` volume
or use the bundled `backend/assets/resume.pdf` as the initial seed. Keep real
`.env` files and application data outside Git.
