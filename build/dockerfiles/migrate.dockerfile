ARG ALPINE_VERSION=3.19.1
ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine as base
RUN go install github.com/rubenv/sql-migrate/...@latest
COPY ./database/migrations /migrations
WORKDIR /migrations
ENTRYPOINT ["sql-migrate", "up", "-env=development"]
