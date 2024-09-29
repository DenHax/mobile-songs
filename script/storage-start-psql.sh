#!/usr/bin/env bash

docker run --name=lib-psql -e POSTGRES_PASSWORD="$POSTGRES_PASSWORD" -p "$POSTGRES_PORT":5432 -d --rm postgres
