FROM node:16-alpine3.11 as build-node

RUN apk --no-cache --virtual build-dependencies add \
        python \
        make \
        g++

# Set the working directory to the Go application
WORKDIR /app

COPY client/ .

RUN npm install

# Build the React app
RUN npm run build

# Use an official Go runtime as a parent image
FROM golang:1.17rc2-alpine3.14 as build-go

WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod .

# Download and install any required Go dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY *.go .

COPY *.txt .

COPY --from=build-node /app ./client

# Build the Go application
RUN go build -o search

EXPOSE 3001

# Start the Go application
ENTRYPOINT ["./search", "-http", ":3001"]