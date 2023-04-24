# Base image
FROM golang:1.20.0

# Set the current working directory inside the container
WORKDIR /app

# Install air
RUN go install github.com/cosmtrek/air@latest

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Clean up the module cache
RUN go mod tidy

# Expose the port
EXPOSE 5000

# Start the application
CMD ["air", "-c", ".air.toml"]
