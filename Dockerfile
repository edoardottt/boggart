# Start from the latest golang base image
FROM golang:latest

# Add env vars
ENV MONGO_CONN="mongodb://mongo:27017"
ENV MONGO_DB="boggart"
ENV SHODAN_KEY=""

# Add Maintainer Info
LABEL maintainer="edoardottt <edoardottt.com>"

# Set the Current Working Directory inside the container
WORKDIR /boggart

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#Build gochanges go file
RUN cd cmd && go build -o boggart

# Run the Go app
CMD ["./cmd/boggart"]

# Honeypot
EXPOSE 8092

# Dashboard
EXPOSE 8093

# API
EXPOSE 8094