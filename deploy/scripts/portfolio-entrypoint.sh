#!/bin/sh
set -eu

if [ -f /data/resume/resume.pdf ]; then
  cp /data/resume/resume.pdf /app/assets/resume.pdf
fi

exec /app/portfolio-api
