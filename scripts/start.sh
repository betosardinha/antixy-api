#!/bin/sh

set -euo pipefail

echo "Running migrations (if any)"

goose up || true

echo "Starting API"

exec air