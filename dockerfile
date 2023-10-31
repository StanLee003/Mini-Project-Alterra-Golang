FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code and go.mod/go.sum files
COPY go.mod .
COPY go.sum .

# Download Go module dependencies
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port that the application will run on
EXPOSE 8080

# Define the command to run the application
CMD ["./main"]
