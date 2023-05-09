FROM alpine:latest

RUN mkdir /app

COPY loggerApp.exe /app

CMD [ "/app/loggerApp.exe"]