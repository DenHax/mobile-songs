#!/usr/bin/env bash

if [ -f ./.env ]; then
  echo "Error: .env file already exists"
else
  cat <<EOL >.env
POSTGRES_PASSWORD=p4ssw0rd
POSTGRES_PORT=5432
POSTGRES_HOST=127.0.0.1
DB_USER=lib-admin
DB_NAME=lib
SSL_MODE=disable
POSTGRES_URL=postgres://\${DB_USER}:\${POSTGRES_PASSWORD}@\${POSTGRES_HOST}:\${POSTGRES_PORT}/\${DB_NAME}?sslmode=\${SSL_MODE}
EOL
  echo ".env file created successfully"
fi

if [ -f ./.env ]; then
  eval "$(grep -v '^#' ./.env | xargs -d '\n' -I {} echo export {})"

  if [ $? -eq 0 ]; then
    echo "Environment activation: succeeded"
  else
    echo "Error: Failed to export environment variables"
  fi
else
  echo "Error: .env file not found"
fi
