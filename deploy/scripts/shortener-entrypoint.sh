#!/bin/sh
set -eu

if [ ! -e /data/storage.db ]; then
  : > /data/storage.db
fi
ln -sf /data/storage.db /app/storage.db

exec /app/shortener
