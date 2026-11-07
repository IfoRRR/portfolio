#!/bin/sh
set -eu

ROOT=/home/copilot/workspace/portfolio
SSH_KEY=/home/copilot/.ssh/github-portfolio

runuser -u copilot -- env HOME=/home/copilot \
  GIT_SSH_COMMAND="ssh -i $SSH_KEY -o IdentitiesOnly=yes" \
  git -C "$ROOT" pull --ff-only

cd "$ROOT/deploy"
docker compose build portfolio-api
docker compose up -d --force-recreate portfolio-frontend portfolio-api nginx

echo "Site updated: http://$(hostname -I | awk '{print $1}'):8081/"
