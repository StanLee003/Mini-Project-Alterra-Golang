# Use an official Golang runtime as a parent image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code and go.mod/go.sum files
COPY go.mod .
COPY go.sum .

# Copy the rest of your application code
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port that the application will run on
EXPOSE 8080

# Define the command to run the application
CMD ["./main"]
