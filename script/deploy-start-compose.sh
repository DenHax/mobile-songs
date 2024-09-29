#!/usr/bin/env bash

docker compose -f ./deployments/compose.yaml up --force-recreate
