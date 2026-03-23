# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

ARG CHANGEME_GOOSE_USER
ARG CHANGEME_GOOSE_PASSWORD
ARG CHANGEME_GOOSE_DBNAME
ARG CHANGEME_GOOSE_HOST
ARG CHANGEME_GOOSE_PORT

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN mkdir -p docs

ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING="user=$CHANGEME_GOOSE_USER password=$CHANGEME_GOOSE_PASSWORD dbname=$CHANGEME_GOOSE_DBNAME sslmode=disable host=$CHANGEME_GOOSE_HOST port=$CHANGEME_GOOSE_PORT"

RUN goose -dir migrations up -allow-missing

RUN swag init

# Build the Go application
RUN go build -o main main.go

# Stage 2: Minimal image for running the app
FROM alpine:latest as runner

#RUN apk add --update supervisor && rm  -rf /tmp/* /var/cache/apk/*

#ADD supervisord.conf /etc/

ENV GO_ENV=production

# Create a directory for the application
WORKDIR /app

#RUN mkdir -p /var/log/supervisor/
#COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN apk add bash

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/resources/bpmn ./resources/bpmn
COPY --from=builder /app/card-pipeline.txt /etc/profile.d
COPY --from=builder /app/card-pipeline.sh /etc/profile.d

# Expose the port on which the application runs (if applicable)
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/bin/bash", "-l", "-c", "./main"]

#CMD ["./main"]
#ENTRYPOINT ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]
