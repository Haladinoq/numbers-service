FROM golang:1.18-alpine

RUN mkdir /app

COPY createChargeService /app
COPY /config/ ./config/

CMD [ "/app/createChargeService","-config","config/config.yaml"]