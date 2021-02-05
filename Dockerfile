FROM golang:1.15.2-alpine3.12

LABEL maintainer="Daniel Lobaton <dlobaton@eklow.ai>"
 
# Set working directory
WORKDIR /app

# Creds
ENV rsUsername=api_score
ENV rsPassword=d627StTYf#y@lzg#Ej1*tmHL

# env vars
ENV GIN_MODE=release
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

# Remove duplicate source files
RUN rm -rf src

# Build the app
RUN go build -o Gotham
# Make port 8080 available to the world outside this container
EXPOSE 8080
# Run the app
CMD ["./Gotham"]
