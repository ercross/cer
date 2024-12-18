FROM golang:1.23-bullseye

# cosmtrek provides auto-reload during development
RUN apt-get update  &&  go install github.com/cosmtrek/air@v1.52.1

# Set working directory
WORKDIR /app

# Create the bin directory with proper permissions
RUN chmod 777 bin/app

EXPOSE 15001