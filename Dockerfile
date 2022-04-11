# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="edoardottt <edoardott@gmail.com>"

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

EXPOSE 8090