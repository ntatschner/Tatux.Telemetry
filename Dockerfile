FROM golang:bullseye

LABEL Author="Nigel Tatschner (ntatschner@gmail.com)"

ENV LISTENONPORT=9000 INFLUXDB_URL='http://localhost' INFLUXDB_PORT=8086 INFLUXDB_ORG='DefaultOrg' INFLUXDB_TOKEN='' INFLUXDB_BUCKET=''

ENV APP_HOME=/app

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

# If a .env file is present, load it
COPY ./.env .env

# load the env file
RUN if [ -f .env ]; then export $(cat .env | xargs); fi

# Copy go mod and sum files
COPY ./src/api/* ${APP_HOME}

RUN go install github.com/ntatschner/Tatux.Telemetry/src/api@latest

# Copy the source code
COPY ./src/api .

RUN go build -o main main.go

# Expose port for the application
EXPOSE $LISTENONPORT

# Command to run the application
CMD ["./main"]