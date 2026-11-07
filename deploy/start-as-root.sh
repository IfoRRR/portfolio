#!/bin/sh
set -eu

cd /home/copilot/workspace/portfolio/deploy

if [ ! -f .env ]; then
  echo "Missing deploy/.env; create it from deploy/.env.example first." >&2
  exit 1
fi

docker compose build
docker compose up -d
docker compose ps

for path in /api/health / /chat/ /shortener/ /quiz/ /board/ /meeting/; do
  ready=0
  for attempt in $(seq 1 30); do
    if curl --fail --silent --show-error "http://127.0.0.1:8081${path}" >/dev/null; then
      echo "OK ${path}"
      ready=1
      break
    fi
    sleep 2
  done
  if [ "$ready" -ne 1 ]; then
    echo "FAILED ${path}" >&2
    docker compose ps
    docker compose logs --tail=100
    exit 1
  fi
done

echo "Portfolio stack is available at http://$(hostname -I | awk '{print $1}'):8081/"
