#!/bin/sh
set -eu

if [ "${SYNC_DEFAULT_RESUME:-0}" = "1" ] && [ -f /app/default-resume/resume.pdf ]; then
  cp /app/default-resume/resume.pdf /data/resume/resume.pdf
elif [ -f /data/resume/resume.pdf ]; then
  cp /data/resume/resume.pdf /app/assets/resume.pdf
fi

if [ -f /data/resume/resume.pdf ]; then
  cp /data/resume/resume.pdf /app/assets/resume.pdf
fi

exec /app/portfolio-api
