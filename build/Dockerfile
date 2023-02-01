# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /api
RUN mkdir /config

COPY . /api
COPY ../config /config


WORKDIR /api

RUN CGO_ENABLED=0 go build -o numbersService ./cmd/numbers-service

RUN chmod +x /api/numbersService

CMD [ "/api/numbersService","-config","/config/config.yml"]