# Effective Mobile service: Songs library

Service start:

```sh
go run cmd/lib/main.go
```

Start with docker compose:
```sh
. ./script/autostart-compose.sh
source .env-compose

make run-compose # if make installed

. ./script/deploy-start-compose.sh # if make not installed
```
