# A Simple Telemetry Collection API - Written in Go(lang)

[![Build API in Docker Container and Deploy Container to Registry](https://github.com/ntatschner/Tatux.Telemetry/actions/workflows/api_build_and_deploy.yml/badge.svg)](https://github.com/ntatschner/Tatux.Telemetry/actions/workflows/api_build_and_deploy.yml)

## Overview
This is a way to collect anonymized metrics about processes - PowerShell functions from a module, CI/CD Pipelines or other automation sources. 
It's is currently in early stages of development, I built this to help be collect usage information about my modules as a way to ensure they are functioning correctly and so I can offer improvements.

At the moment I'm implementing a influxdb v2 database backend - NOT INCLUDED

All the code to is shared freely, I will also link the modules I'm embedding this in to so you can maybe utilise this for yourself. 

## Running the Docker Image

You can find the Docker image on Docker Hub [here](https://hub.docker.com/r/tatux/telemetry).

To run the Docker image, use the following command:

```bash
docker run -e LISTENONPORT=9000 -e INFLUXDB_URL=http://localhost -e INFLUXDB_PORT=8086 -e INFLUXDB_ORG=DefaultOrg -e INFLUXDB_TOKEN=my-token -e INFLUXDB_BUCKET=my-bucket -p 9000:9000 tatux/telemetry:latest
```

In this command, `LISTENONPORT`, `INFLUXDB_URL`, `INFLUXDB_PORT`, `INFLUXDB_ORG`, `INFLUXDB_TOKEN`, and `INFLUXDB_BUCKET` are environment variables that are being passed to the Docker container. Replace `http://localhost`, `8086`, `DefaultOrg`, `my-token`, and `my-bucket` with your actual values.

The `-p 9000:9000` option maps port 9000 in the Docker container to port 9000 on your host machine, so you can access your application at `http://localhost:9000`.

Docker Compose Example:
```yaml
  telemetry:
    image: tatux/telemetry:latest
    container_name: telemetry
    restart: always
    ports:
    - "9000:9000"
    environment:
      LISTENONPORT: 9000
      INFLUXDB_UR: http://localhost
      INFLUXDB_PORT: 8086
      INFLUXDB_TOKEN: YOUR_TOKEN
      INFLUXDB_BUCKET: YOUR_BUCKET
      INFLUXDB_ORG: YOUR_ORG
```
## Dockerfile Breakdown

Here's a breakdown of what each part of the Dockerfile does:

- `FROM golang:latest as builder`: This line starts a new build stage with the `golang:latest` image as the base. The `as builder` part names the build stage "builder", so it can be referred to later.

- `ENV LISTENONPORT=9000 INFLUXDB_URL='http://localhost' INFLUXDB_PORT=8086 INFLUXDB_ORG='DefaultOrg' INFLUXDB_TOKEN='' INFLUXDB_BUCKET=''`: This line sets environment variables that the application uses.

- `RUN mkdir -p builder`: This line creates a new directory named "builder".

- `WORKDIR /builder`: This line changes the working directory to "/builder".

- `COPY ./src/api/ /builder`: This line copies the application code from the "src/api" directory in your host machine to the "/builder" directory in the image.

- `RUN go mod download`: This line downloads the Go modules needed by the application.

- `RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .`: This line builds the Go application and produces an executable named "main".

- `FROM alpine:latest`: This line starts a new build stage with the `alpine:latest` image as the base.

- `RUN apk --no-cache add ca-certificates`: This line installs the ca-certificates package, which is needed to make HTTPS requests.

- `ENV APP_HOME=/app`: This line sets an environment variable named "APP_HOME".

- `RUN mkdir -p "$APP_HOME"`: This line creates a new directory for the application.

- `WORKDIR "$APP_HOME"`: This line changes the working directory to the application directory.

- `COPY --from=builder /builder/main .`: This line copies the "main" executable from the "builder" build stage to the current directory in the image.

- `EXPOSE $LISTENONPORT`: This line informs Docker that the container listens on the specified network port at runtime.

- `CMD ["./main"]`: This line specifies the command to run when the container starts.

Again, this is in the early stages and I would use any code with a pitch of salt. :)