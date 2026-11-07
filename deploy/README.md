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

The portfolio frontend is a separate `portfolio-frontend` service. To update
only Vue files without rebuilding the API or other applications:

```bash
docker compose build portfolio-frontend
docker compose up -d --no-deps portfolio-frontend
```

The nginx container keeps its existing configuration and routes `/api/...` to
the API while serving `/` from the frontend service.

As a one-command root deployment from the prepared VPS workspace:

```bash
/home/copilot/workspace/portfolio/deploy/start-as-root.sh
```

The script builds the images, starts the stack, prints container status, and
checks the public routes through nginx. It does not modify VPN configuration.

The current compose stack deploys the portfolio, WebSocket chat, shortener,
quiz, Kanban board and meeting-room MVP. The meeting room is available at
`/meeting/` and uses an internal PostgreSQL service; neither PostgreSQL nor the
application port is published to the VPS.

For the first deployment, copy the resume into the `portfolio-resume` volume
or use the bundled `backend/assets/resume.pdf` as the initial seed. Keep real
`.env` files and application data outside Git.
