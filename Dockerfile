# Start from golang base image
FROM golang:1.16-alpine

ENV LISTENONPORT=9000 INFLUXDB_URL='http://localhost' INFLUXDB_PORT=8086 INFLUXDB_ORG='DefaultOrg' INFLUXDB_TOKEN='' INFLUXDB_BUCKET=''

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./src/go.mod ./src/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY ./src .

# Build the application
RUN go build -o main .

# Expose port for the application
EXPOSE $LISTENONPORT

# Command to run the application
CMD ["/app/main"]