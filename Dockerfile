FROM golang:1.21 AS builder
LABEL maintainer="Timam <mail@timam.io>"

WORKDIR /pokemon
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o api ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /pokemon/api .

EXPOSE 8080
CMD ["./api"]
