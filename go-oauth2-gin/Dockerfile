# FROM golang:1.24.4-alpine3.22

# WORKDIR /app

# COPY . .

# RUN go get github.com/lib/pq

# ENV GOOS=linux
# ENV GOARCH=amd64

# RUN go mod tidy

# RUN go build -o api .

# EXPOSE 8080

# CMD [ "./api" ]


# used for Outh2 project
# Use official Go image with Alpine for small size
# FROM golang:1.24.4-alpine3.22

# Set working directory
# WORKDIR /app

# Copy go.mod and go.sum first (to cache dependencies)
# COPY go.mod ./
# COPY go.sum ./

# Download dependencies
# RUN go mod download

# Copy the full source code
# COPY . .

# # Debug log: what's in the folder
# RUN echo "🧱 Files before building:" && ls -la

# Cross-compile for Linux (important!)
# ENV GOOS=linux
# ENV GOARCH=amd64

# Build Go binary
# RUN go build -o api .

# # Debug log: is binary built?
# RUN echo "🔍 After build:" && ls -la

# Expose port
# EXPOSE 8080

# Run the binary
# CMD ["air"]

# Use official Go image with Alpine
FROM golang:1.24.4-alpine3.22

# Install necessary packages and Air
RUN apk add --no-cache git curl && \
    go install github.com/air-verse/air@latest

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the full source code, including Air config
COPY . .

# Expose the app port
EXPOSE 8080

# Start with Air (no need to build manually)
CMD ["air"]
