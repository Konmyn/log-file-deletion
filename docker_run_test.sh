#!/usr/bin/env bash

export $(cat .env | xargs)
docker run --rm -it -v /tmp:/logs "$image" /app/log-walk --work-hour=13 --path=/logs --preserve-hour=1 --nap-time=50 --dry-run
