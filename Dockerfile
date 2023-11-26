# Use the official Go image as the base image
FROM golang:1.20.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files into the container
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the Go application
RUN go build -o orders-api .

# Expose port 3000 for the application
EXPOSE 3000

# Run the compiled binary when the container starts
CMD ["./orders-api"]
