# Production deployment

The compose file expects these repositories as sibling directories:

```text
/home/ubuntu/
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

The portfolio frontend is built by
`dockerfiles/portfolio-frontend.Dockerfile` from the current Git checkout.
After frontend changes, rebuild `portfolio-frontend` and recreate it:

```bash
docker compose build portfolio-frontend
docker compose up -d --force-recreate portfolio-frontend nginx
```

Nginx routes `/api/...` to the API while serving `/` from the frontend
service.

As a one-command root deployment from the prepared VPS workspace:

```bash
/home/copilot/workspace/portfolio/deploy/start-as-root.sh
```

For regular site updates, including the frontend and resume:

```bash
/home/copilot/workspace/portfolio/deploy/update-site.sh
```

The script builds the images, starts the stack, prints container status, and
checks the public routes through nginx. It does not modify VPN configuration.

The current compose stack deploys the portfolio, WebSocket chat, shortener,
quiz, Kanban board and meeting-room MVP. The meeting room is available at
`/meeting/` and uses an internal PostgreSQL service; neither PostgreSQL nor the
application port is published to the VPS.

For the first deployment, use the bundled `backend/assets/resume.pdf` as the
initial seed. Keep real `.env` files and application data outside Git.

## Recovery after a VPS reset

The deployment code is stored in GitHub, but `.env` and Docker volumes are
not. Keep a backup directory outside the VPS. Create one before a migration:

```bash
sudo sh /home/ubuntu/portfolio/deploy/scripts/backup.sh
```

After provisioning a new Ubuntu VPS with Git, Docker and Compose, run:

```bash
sudo sh /home/ubuntu/portfolio/deploy/scripts/restore-vps.sh
```

The restore script clones the portfolio and its sibling repositories, creates a
new database password only when `deploy/.env` does not exist, validates Compose,
builds the images and starts the stack. To restore old data, copy the backed-up
`.env` into `deploy/.env` before starting and import the volume archives from
the backup directory. Never commit `.env` or upload it to a public repository.
