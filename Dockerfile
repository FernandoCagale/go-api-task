FROM golang:1.9

RUN mkdir -p /app

WORKDIR /app

ADD build/api /app/api

ENTRYPOINT ["./api"]