#!/bin/sh
set -eu

ROOT=/home/copilot/workspace/portfolio
NODE_BIN=${NODE_BIN:-/home/copilot/node/bin}
BUILD_DIR=$(mktemp -d)
trap 'rm -rf "$BUILD_DIR"' EXIT

export PATH="$NODE_BIN:$PATH"

if [ ! -x "$NODE_BIN/npm" ]; then
  echo "Node.js is missing at $NODE_BIN. Install it once or set NODE_BIN." >&2
  exit 1
fi

cp "$ROOT/frontend/package.json" "$ROOT/frontend/package-lock.json" "$BUILD_DIR/"
cp "$ROOT/frontend/vite.config.js" "$BUILD_DIR/"
cp -R "$ROOT/frontend/src" "$ROOT/frontend/index.html" "$BUILD_DIR/"
cd "$BUILD_DIR"
npm ci

npm run build
mkdir -p "$ROOT/frontend/dist/assets"
rm -f "$ROOT/frontend/dist/index.html"
rm -rf "$ROOT/frontend/dist/assets"/*
cp "$BUILD_DIR/dist/index.html" "$ROOT/frontend/dist/index.html"
cp -R "$BUILD_DIR/dist/assets/." "$ROOT/frontend/dist/assets/"
echo "Frontend build updated: $ROOT/frontend/dist"
