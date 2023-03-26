FROM golang:1.20.1-alpine3.17 AS builder
COPY ./ /badger-db
RUN rm -rf /badger-db/.env && touch /badger-db/.env
WORKDIR /badger-db
RUN go build

FROM alpine:3.17
COPY --from=builder /badger-db/badger-db /badger-db/badger-db
COPY --from=builder /badger-db/.env /badger-db/.env
RUN mkdir /badger-db/data
RUN chown 1000:1000 /badger-db/data
USER 1000
WORKDIR /badger-db
CMD ["./badger-db"]
