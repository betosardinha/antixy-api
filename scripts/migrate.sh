#!/bin/sh

set -euo pipefail

COMPOSE_FILE="docker-compose.dev.yml"
CMD="${@:-up}"

echo "Running migrations: $CMD"

docker compose -f $COMPOSE_FILE run --rm migrate goose $CMD
