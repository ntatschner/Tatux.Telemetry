FROM golang:latest as builder

LABEL \
    org.opencontainers.image.authors="ntatschner@gmail.com" \
    org.opencontainers.image.created=$CREATED \
    org.opencontainers.image.version=$VERSION \
    org.opencontainers.image.revision=$COMMIT \
    org.opencontainers.image.url="https://github.com/ntatschner/Tatux.Telemetry" \
    org.opencontainers.image.documentation="https://github.com/ntatschner/Tatux.Telemetry" \
    org.opencontainers.image.source="https://github.com/ntatschner/Tatux.Telemetry" \
    org.opencontainers.image.title="A Telemetry Collection API" \
    org.opencontainers.image.description="A Telemetry Collection API that sends data to InfluxDB."

LABEL Author="Nigel Tatschner (ntatschner@gmail.com)"

ENV LISTENONPORT=9000 INFLUXDB_URL='http://localhost' INFLUXDB_PORT=8086 INFLUXDB_ORG='DefaultOrg' INFLUXDB_TOKEN='' INFLUXDB_BUCKET='' TRUSTEDPROXIES='' IP2LOCATION_API_KEY=''

RUN mkdir -p builder

WORKDIR /builder

COPY ./src/api/ /builder

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV APP_HOME=/app

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY --from=builder /builder/main .

EXPOSE $LISTENONPORT

CMD ["./main"]
