# Stage 1: Build the Go application
FROM golang:1.20.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o orders-api .
# Stage 2: Create a minimal image
FROM alpine:3.15 as root-certs
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app
# Stage 3: Create the final image
FROM scratch
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/orders-api /orders-api
USER app
# Expose port 3000 for the application
EXPOSE 3000
# Run the compiled binary when the container starts
CMD ["/orders-api"]