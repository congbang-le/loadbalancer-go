# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app
COPY go.mod ./
COPY *.go ./

RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]