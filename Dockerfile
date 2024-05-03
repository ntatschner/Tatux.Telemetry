FROM golang:latest as builder

LABEL Author="Nigel Tatschner (ntatschner@gmail.com)"

ENV LISTENONPORT=9000 INFLUXDB_URL='http://localhost' INFLUXDB_PORT=8086 INFLUXDB_ORG='DefaultOrg' INFLUXDB_TOKEN='' INFLUXDB_BUCKET=''

ENV APP_HOME=/app

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY ./src/api/* ${APP_HOME}

RUN go get github.com/ntatschner/Tatux.Telemetry/src/api/handlers \
    && go get github.com/ntatschner/Tatux.Telemetry/src/api/app \
    && go get github.com/ntatschner/Tatux.Telemetry/src/api/system

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"\main .

EXPOSE $LISTENONPORT

CMD ["./main"]