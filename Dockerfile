# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .


# Set the working directory for the build context
WORKDIR /app/cmd/fileshare


# Build the Go application
RUN go build -o /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
