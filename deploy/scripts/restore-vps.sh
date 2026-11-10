#!/bin/sh
set -eu

ROOT=${PORTFOLIO_ROOT:-/home/ubuntu/portfolio}
PARENT=$(dirname "$ROOT")
ENV_FILE="$ROOT/deploy/.env"

clone_if_missing() {
  repo=$1
  directory=$2

  if [ -d "$PARENT/$directory/.git" ]; then
    runuser -u ubuntu -- git -C "$PARENT/$directory" pull --ff-only
  else
    runuser -u ubuntu -- git clone "$repo" "$PARENT/$directory"
  fi
}

if [ "$(id -u)" -ne 0 ]; then
  echo "Run this script as root." >&2
  exit 1
fi

mkdir -p "$PARENT"
if [ ! -d "$ROOT/.git" ]; then
  runuser -u ubuntu -- git clone https://github.com/IfoRRR/portfolio.git "$ROOT"
else
  runuser -u ubuntu -- git -C "$ROOT" pull --ff-only
fi

clone_if_missing https://github.com/IfoRRR/WEB-socket.git websocket
clone_if_missing https://github.com/IfoRRR/URL-SHORTENER.git url-shortener
clone_if_missing https://github.com/IfoRRR/vk-quiz.git vk-quiz
clone_if_missing https://github.com/IfoRRR/vk-kanban.git vk-kanban
clone_if_missing https://github.com/IfoRRR/meeting-room-mvp.git meeting-room-mvp

if [ ! -f "$ENV_FILE" ]; then
  umask 077
  printf 'MEETING_DB_PASSWORD=%s\n' "$(openssl rand -hex 32)" > "$ENV_FILE"
  chown ubuntu:ubuntu "$ENV_FILE"
fi

cd "$ROOT"
docker compose --env-file "$ENV_FILE" -f deploy/docker-compose.yml config >/dev/null
docker compose --env-file "$ENV_FILE" -f deploy/docker-compose.yml build
docker compose --env-file "$ENV_FILE" -f deploy/docker-compose.yml up -d
docker compose --env-file "$ENV_FILE" -f deploy/docker-compose.yml ps

echo "Restore complete: http://$(hostname -I | awk '{print $1}'):8081/"
