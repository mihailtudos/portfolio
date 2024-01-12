# The base go-image
FROM golang:1.21.6-bullseye
 
# Create a directory for the app
RUN mkdir /app

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies separately
COPY go.mod .
COPY go.sum .

# Download dependencies (if any)
RUN go mod tidy

# Install fresh
RUN go install github.com/cosmtrek/air@latest

# Copy the entire project to the app directory
COPY . .

# Run the server using fresh
CMD ["air"]