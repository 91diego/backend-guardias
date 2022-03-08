FROM golang:1.16.3-buster as build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/backend-guardias main.go

FROM debian:buster-slim

COPY --from=build /app/bin/backend-guardias /usr/local/bin/backend-guardias
COPY wait-for-it.sh /usr/local/bin/wait-for-it
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint

ENTRYPOINT ["docker-entrypoint"]
CMD ["backend-guardias"]
