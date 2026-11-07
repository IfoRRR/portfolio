#!/bin/sh
set -eu

if [ -f /data/resume/resume.pdf ]; then
  cp /data/resume/resume.pdf /app/assets/resume.pdf
elif [ -f /app/default-resume/resume.pdf ]; then
  cp /app/default-resume/resume.pdf /data/resume/resume.pdf
fi

exec /app/portfolio-api
