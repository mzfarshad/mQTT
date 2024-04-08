# Use official golang image as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /mqtt

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Set up environment variables from build arguments
ARG DB_HOST
ARG DB_USER
ARG DB_PASS
ARG DB_NAME
ARG DB_PORT
ARG DB_TIMEZONE

ENV DB_HOST=$DB_HOST
ENV DB_USER=$DB_USER
ENV DB_PASS=$DB_PASS
ENV DB_NAME=$DB_NAME
ENV DB_PORT=$DB_PORT
ENV DB_TIMEZONE=$DB_TIMEZONE

# Expose the port your application runs on

# Command to run the executable
CMD ["./main"]
