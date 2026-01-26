# --- Build Stage ---
# Use an official Go image as the build stage base
FROM golang:1.25-alpine AS builder

# Install make
RUN apk update && apk add --no-cache make

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to leverage Docker's build cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build from Makefile
RUN make 

# --- Final Stage ---
# Start from a minimal, non-golang base image (e.g., 'scratch' or 'alpine')
FROM alpine:latest

# Set the working directory for the final application
WORKDIR /app

# Copy the compiled binary from the 'builder' stage to the final image
COPY --from=builder /app/build/vidpovid-bot-go /app/vidpovid-bot-go
COPY --from=builder /app/build/.env.example /app/.env

# Command to run the application when the container starts
CMD ["/app/vidpovid-bot-go"]
