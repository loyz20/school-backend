#!/bin/sh
set -e

echo "📦 Running migrations..."

migrate -path /migrations \
  -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable" \
  -verbose up

echo "✅ Migration finished."
