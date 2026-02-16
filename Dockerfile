FROM golang:1.24.2 AS builder

WORKDIR /url_shortener

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build main.go 

ENV STORAGE_TYPE=memory
ENV DB_CONN_STR="postgres://postgres:12345@localhost/postgres?sslmode=disable"

CMD [ "./main" ]