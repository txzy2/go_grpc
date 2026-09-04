#!/bin/bash
set -e

if [ ! -f .env ]; then
  cp .env.example .env
  echo ".env создан из .env.example"
fi

docker compose -f compose.dev.yml up -d
echo "Сервисы запущены"
