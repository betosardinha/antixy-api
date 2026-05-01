#!/bin/sh

set -euo pipefail

COMPOSE_FILE="docker-compose.dev.yml"
CMD=${1:-up}

echo "Running migrations: $CMD"

docker compose -f $COMPOSE_FILE run --rm migrate \
  sh -c "goose -dir internal/database/migrations postgres \"\$DATABASE_URL\" $CMD"
