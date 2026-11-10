#!/bin/sh
set -eu

ROOT=${PORTFOLIO_ROOT:-/home/ubuntu/portfolio}
COMPOSE="$ROOT/deploy/docker-compose.yml"
ENV_FILE="$ROOT/deploy/.env"
BACKUP_DIR=${1:-/home/ubuntu/backups/portfolio-$(date -u +%Y%m%d-%H%M%S)}

if [ "$(id -u)" -ne 0 ]; then
  echo "Run this script as root." >&2
  exit 1
fi

if [ ! -f "$ENV_FILE" ]; then
  echo "Missing $ENV_FILE" >&2
  exit 1
fi

mkdir -p "$BACKUP_DIR/config" "$BACKUP_DIR/volumes"
cp "$ENV_FILE" "$BACKUP_DIR/config/.env"
cp "$ROOT/deploy/docker-compose.yml" "$BACKUP_DIR/config/docker-compose.yml"
cp "$ROOT/deploy/nginx/nginx.conf" "$BACKUP_DIR/config/nginx.conf"
cp "$ROOT/deploy/nginx/portfolio-frontend.conf" "$BACKUP_DIR/config/portfolio-frontend.conf"
cp "$ROOT/backend/assets/resume.pdf" "$BACKUP_DIR/config/resume.pdf"

docker compose --env-file "$ENV_FILE" -f "$COMPOSE" ps > "$BACKUP_DIR/config/containers.txt"
ss -tulpn > "$BACKUP_DIR/config/ports.txt"

for logical_volume in $(docker compose --env-file "$ENV_FILE" -f "$COMPOSE" config --volumes); do
  volume="deploy_${logical_volume}"
  if ! docker volume inspect "$volume" >/dev/null 2>&1; then
    volume="$logical_volume"
  fi
  if docker volume inspect "$volume" >/dev/null 2>&1; then
    docker run --rm \
      -v "$volume:/source:ro" \
      -v "$BACKUP_DIR/volumes:/backup" \
      alpine:3.22 \
      tar czf "/backup/${logical_volume}.tar.gz" -C /source .
  fi
done

chmod -R go-rwx "$BACKUP_DIR"
echo "Backup created at $BACKUP_DIR"
echo "Copy this directory to storage outside the VPS."
