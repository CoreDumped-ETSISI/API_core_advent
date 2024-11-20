# Use an Alpine-based Go image
FROM golang:alpine

# Install libc6-compat to handle potential library compatibility issues
RUN apk add --no-cache build-base libc6-compat

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go modules and sum files, as well as the .env file
COPY go.mod go.sum .env ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app for Linux and amd64 architecture
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /app/main .


# List the files to confirm if /app/main exists
RUN ls -l /app

# Expose port
EXPOSE 8000
RUN chmod +x /app/main

# Run the Go app
CMD [ "/app/main" ]
