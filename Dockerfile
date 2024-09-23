# Base image for building the frontend (SvelteKit app)
FROM node:18 as client-build

# Set the working directory to /client
WORKDIR /client

# Copy the package.json and package-lock.json to install dependencies
COPY client/package*.json ./

# Install dependencies for the frontend
RUN npm install

# Copy the rest of the client app code to the container
COPY client/ .

# Build the SvelteKit app for production
RUN npm run build

# ----

# Base image for building the backend (Go API)
FROM golang:1.23-alpine as server-build

# Set the working directory to /server
WORKDIR /server

# Copy the Go module files
COPY server/go.mod server/go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the server app code
COPY server/ .

# Build the Go application for production
RUN go build -o /server/app

# ----

# Base image for the Python scraper
FROM python:3.11-alpine as python-scraper

# Set the working directory to /scraper
WORKDIR /scraper

# Copy the Python scraper code
COPY scraper/ .

# Production image (final stage)
FROM alpine:3.18

# Install necessary packages to run the Go server and serve static files
RUN apk --no-cache add ca-certificates && \
    apk add --no-cache python3 py3-pip

WORKDIR /
COPY data.json .

# Set the working directory for the final container
WORKDIR /app

# Copy the built Go backend
COPY --from=server-build /server/app /app/app

# Copy the built SvelteKit frontend to /app/public for static serving
COPY --from=client-build /client/build /app/public

# Copy the Python scraper
COPY --from=python-scraper /scraper /app/scraper

# Install Python dependencies
RUN pip install -r /app/scraper/requirements.txt

# Expose the port (adjust if necessary)
EXPOSE 8080

# Command to run the Go server
CMD ["/app/app"]
