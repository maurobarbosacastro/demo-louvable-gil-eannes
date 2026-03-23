# First stage: build the Go binary
FROM golang:1.22-alpine  AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

ENV GO111MODULE=on

# Build the binary with additional build flags
RUN go build -o app


# Second stage: create the final image
FROM alpine:3.14.10
WORKDIR /app

RUN apk add bash

ENV GIN_MODE=release
ENV PORT=8080

COPY --from=builder /app/card-pipeline.txt /etc/profile.d
COPY --from=builder /app/card-pipeline.sh /etc/profile.d

COPY --from=builder /app/app .
# COPY --from=builder /app/.env . #UNCOMENT THIS LINE ONLY TO RUN IN LOCAL MACHINE

RUN echo "alpine:3.14.10" > /app/version.txt

RUN apk add --no-cache tzdata
ENV TZ=Europe/Lisbon

EXPOSE 8080

ENTRYPOINT ["/bin/bash", "-l", "-c", "./app"]

VOLUME ["/app/images"]
