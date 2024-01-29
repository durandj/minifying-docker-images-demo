FROM golang:1.21 AS build

WORKDIR /code

COPY . .
RUN CGO_ENABLED=0 go build -o service cmd/main.go


FROM debian:bookworm

WORKDIR /srv

COPY --from=build /code/service ./service

CMD ["/srv/service"]
