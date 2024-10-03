#!/usr/bin/env bash

set -e

shift
cmd="$@"

if [[ -z "$POSTGRES_PASSWORD" || -z "$POSTGRES_HOST" || -z "$DB_NAME" ]]; then
  echo "Error: Required environment variables are not set."
  echo "Please set POSTGRES_PASSWORD, POSTGRES_HOST, and DB_NAME."
  exit 1
fi

until PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -U "$DB_NAME" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd
