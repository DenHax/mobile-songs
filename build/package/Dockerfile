FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache bash git make gcc gettext musl-dev

COPY ./go.mod ./go.sum .

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/lib ./cmd/lib/main.go

FROM alpine:3.20 AS runner

COPY --from=builder /usr/local/src/bin/lib /

RUN apk add --no-cache ca-certificates

COPY ./configs/config.yaml /config.yaml

ENV CONFIG_PATH  /config.yaml

CMD ["/lib"]
