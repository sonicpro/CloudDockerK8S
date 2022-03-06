# Download Go toolchain as base image
FROM golang:1.17

# Set convenient work dir
WORKDIR /app

# Copy external source code file to the container
COPY main.go ./

# Compile
RUN go build -o hello-go main.go

# Container entry point
CMD ["./hello-go"]
