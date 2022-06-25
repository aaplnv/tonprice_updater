FROM golang:1.18-buster as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o main
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /build/main /app/
WORKDIR /app
COPY settings.env ./
CMD ["./main"]