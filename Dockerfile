FROM golang:1.13.5-buster as builder

WORKDIR /app
COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o mikrotik-exporter_linux_amd64 . && \
    chmod 775 mikrotik-exporter_linux_amd64

# Production container
FROM debian:buster-slim

EXPOSE 9436
WORKDIR /app
COPY scripts/start.sh .
COPY --from=builder /app/mikrotik-exporter_linux_amd64 ./mikrotik-exporter

ENTRYPOINT ["/app/start.sh"]
