#!/bin/sh

set -e

echo "Checking migrations..."

if goose -dir /app/migrations postgres "$DATABASE_URL" status | grep -q "Pending"; then
  echo "Running migrations..."
  goose -dir /app/migrations postgres "$DATABASE_URL" up
else
  echo "No migrations needed"
fi

echo "Starting API..."

air
