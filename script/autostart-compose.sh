#!/usr/bin/env bash

if [ -f ./.env-compose ]; then
  echo "Error: .env file already exists"
else
  cat <<EOL >.env-compose
POSTGRES_PASSWORD=p4ssw0rd
POSTGRES_PORT=5432
POSTGRES_HOST=storage
DB_USER=song-admin
DB_NAME=songs
SSL_MODE=disable
POSTGRES_URL=postgres://\${DB_USER}:\${POSTGRES_PASSWORD}@\${POSTGRES_HOST}:\${POSTGRES_PORT}/\${DB_NAME}?sslmode=\${SSL_MODE}
EOL
  echo ".env-compose file created successfully"
fi

if [ -f ./.env-compose ]; then
  eval "$(grep -v '^#' ./.env-compose | xargs -d '\n' -I {} echo export {})"

  if [ $? -eq 0 ]; then
    echo "Environment activation: succeeded"
  else
    echo "Error: Failed to export environment variables"
  fi
else
  echo "Error: .env file not found"
fi
