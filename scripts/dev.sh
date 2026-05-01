#!/bin/sh

set -euo pipefail

COMPOSE_FILE="docker-compose.dev.yml"
CMD=${1-}

echo "Starting dev environment"

docker compose -f docker-compose.dev.yml up \
  -d --build \
  db redis api \
  $CMD