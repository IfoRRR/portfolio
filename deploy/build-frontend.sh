#!/bin/sh
set -eu

ROOT=/home/copilot/workspace/portfolio
NODE_BIN=${NODE_BIN:-/home/copilot/node/bin}

export PATH="$NODE_BIN:$PATH"
cd "$ROOT/frontend"

if [ ! -x "$NODE_BIN/npm" ]; then
  echo "Node.js is missing at $NODE_BIN. Install it once or set NODE_BIN." >&2
  exit 1
fi

if [ ! -x node_modules/.bin/vite ]; then
  rm -rf node_modules
  npm ci
fi

npm run build
echo "Frontend build updated: $ROOT/frontend/dist"
