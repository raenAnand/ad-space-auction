# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o auction

# Expose the port on which the application will run
EXPOSE 8001

# Define the command to run the application
CMD ["./auction"]
