FROM golang:1.15.2-alpine3.12

LABEL maintainer="Daniel Lobaton <dlobaton@eklow.ai>"

ENV GIN_MODE=release
 
# Set working directory
WORKDIR /app

# DB env - dev
ENV dbHost=gotham-dev.cuevwe5bpzjq.us-east-2.rds.amazonaws.com
ENV dbPort=5432
ENV dbName=gotham
ENV dbUser=Quiver
ENV dbPassword=JC7lj4odvCQcEFVJoUo2

# Download dependencies
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

# Copy all the app sources (recursively copies files and directories from the host into the image)
COPY src/ .

# Build the app
RUN go build -o Gotham

# Remove source files
RUN find . -name "*.go" -type f -delete 
RUN find . -name "go.*" -type f -delete
RUN rm -rf src
# Make port 5000 available to the world outside this container
EXPOSE 8080

# Run the app
CMD ["./Gotham"]
