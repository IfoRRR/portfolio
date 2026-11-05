# Production deployment

The compose file expects these repositories as sibling directories:

```text
/opt/
├── portfolio/
├── vk-quiz/
├── vk-kanban/
├── websocket/
└── url-shortener/
```

Run from this directory:

```bash
docker compose build
docker compose up -d
```

Only nginx publishes `8081:80`. Application ports remain inside the Docker network.
