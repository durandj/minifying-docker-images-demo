FROM golang:1.21 AS build

WORKDIR /code

COPY . .
RUN CGO_ENABLED=0 go build -o service cmd/main.go


FROM gcr.io/distroless/static-debian11:nonroot

WORKDIR /srv

COPY --from=build /code/service /srv/service

CMD ["/srv/service"]
