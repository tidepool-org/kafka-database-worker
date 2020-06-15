# Development
FROM golang:1.14-alpine AS builder
WORKDIR /go/src/github.com/tidepool-org/kafka-database-worker

RUN adduser -D tidepool && \
    chown -R tidepool /go/src/github.com/tidepool-org/kafka-database-worker
USER tidepool
COPY --chown=tidepool . .
RUN ./build.sh
CMD ["./dist/kafka-database-worker"]
