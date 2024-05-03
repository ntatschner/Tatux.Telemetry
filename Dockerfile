FROM golang:latest as builder

LABEL Author="Nigel Tatschner (ntatschner@gmail.com)"

ENV LISTENONPORT=9000 INFLUXDB_URL='http://localhost' INFLUXDB_PORT=8086 INFLUXDB_ORG='DefaultOrg' INFLUXDB_TOKEN='' INFLUXDB_BUCKET=''


RUN mkdir -p builder

WORKDIR /builder

COPY ./src/api/ /builder

#RUN go get github.com/ntatschner/Tatux.Telemetry/src/api

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