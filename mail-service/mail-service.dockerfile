FROM alpine:latest

RUN mkdir /app

COPY mailerApp.exe /app

CMD [ "/app/mailerApp.exe"]