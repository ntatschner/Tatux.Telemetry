# Start from golang base image
FROM golang:1.16-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port for the application
EXPOSE 1323

# Command to run the application
CMD ["/app/main"]