#!/bin/sh
set -eu

if [ ! -e /data/messages.db ]; then
  : > /data/messages.db
fi
ln -sf /data/messages.db /app/messages.db

exec /app/websocket
