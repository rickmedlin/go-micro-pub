FROM alpine:latest

RUN mkdir /app

COPY frontEndApp.exe /app

CMD [ "/app/frontEndApp.exe"]