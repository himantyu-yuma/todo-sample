FROM golang:1.22.4

WORKDIR /go/src
COPY . .
ENV GO111MODULE=on

RUN apt-get update

RUN go get -u github.com/air-verse/air && \
    go build -o /go/bin/air github.com/air-verse/air
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go get -u github.com/swaggo/swag/cmd/swag && \
    go install github.com/swaggo/swag/cmd/swag@latest

CMD ["air", "-c", ".air.toml"]
