FROM golang:1.21

WORKDIR /code

COPY . .
RUN CGO_ENABLED=0 go build -o service cmd/main.go

CMD ["/code/service"]
