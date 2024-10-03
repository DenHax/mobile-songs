FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache bash git make gcc gettext musl-dev

COPY ./go.mod ./go.sum .

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/song ./cmd/lib/main.go

FROM alpine:3.20 AS runner

RUN apk add --no-cache ca-certificates postgresql-client

COPY ./script/wait-storage.sh /

RUN chmod +x wait-storage.sh

COPY ./configs/config.yaml /config.yaml

ENV CONFIG_PATH  /config.yaml

COPY ./migrations /migrations

COPY --from=builder /usr/local/src/bin/song /

CMD ["/song"]
