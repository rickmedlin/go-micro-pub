SHELL=cmd.exe
BROKER_BINARY=brokerApp.exe
AUTH_BINARY=authApp.exe
LOGGER_BINARY=loggerServiceApp.exe
MAIL_BINARY=mailerApp.exe
LISTENER_BINARY=listenerApp.exe
FRONT_BINARY=frontEndApp.exe

## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_logger build_mail build_listener
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

## build_front_linux: builds the front end binary as a linux executable
build_front_linux:
	@echo Building front end linux binary...
	chdir ..\front-end && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${FRONT_BINARY} ./cmd/web
	@echo Done!

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo Building broker binary...
	chdir ..\broker-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BROKER_BINARY} ./cmd/api
	@echo Done!

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo Building logger binary...
	chdir ..\logger-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${LOGGER_BINARY} ./cmd/api
	@echo Done!

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo Building listener binary...
	chdir ..\listener-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${LISTENER_BINARY} .
	@echo Done!

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo Building auth binary...
	chdir ..\authentication-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${AUTH_BINARY} ./cmd/api
	@echo Done!

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo Building mail binary...
	chdir ..\mail-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${MAIL_BINARY} ./cmd/api
	@echo Done!

## start: starts the front end
start: build_front
	@echo Starting front end
	chdir ..\front-end && start /B ${FRONT_BINARY} &

## stop: stop the front end
stop:
	@echo Stopping front end...
	@taskkill /IM "${FRONT_BINARY}" /F
	@echo "Stopped front end!"