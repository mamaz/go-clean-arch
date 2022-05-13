# Build stage
FROM golang:1.16-alpine AS builder

# Set ARGs
ARG APP_NAME=go-clean-arch

# Set workdir
WORKDIR /app

# Copy all project code
ADD . .

# Download dependencies
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o /tmp/app main.go

# Final stage
FROM alpine:latest AS production

# Copy output binary file from build stage
COPY --from=builder /tmp/app .

# Run the app
CMD ["./app"]