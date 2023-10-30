# Use an official Go runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the dependencies for caching
COPY go.mod go.sum ./

#Run to download dependencies
RUN go mod download

# Copy the local source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8080

# Run the application
CMD ["./main"]
