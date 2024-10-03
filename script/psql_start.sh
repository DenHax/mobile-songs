#!/usr/bin/env bash

docker run --name=song-psql-serv \
  -e POSTGRES_USER="$DB_USER" \
  -e POSTGRES_PASSWORD="$POSTGRES_PASSWORD" \
  -e POSTGRES_DB=DB_NAME \
  -p "$POSTGRES_PORT":5432 \
  postgres:16-alpine
