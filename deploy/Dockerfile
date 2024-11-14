FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk update && \
    apk add build-base

RUN go mod download

COPY ./ ./

RUN go build -o app /app/cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/ /app/

CMD [ "./app" ]