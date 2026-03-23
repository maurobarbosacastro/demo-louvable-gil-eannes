# Base docker image --> Alpine
FROM alpine:latest
RUN command echo 'http://dl-cdn.alpinelinux.org/alpine/v3.9/main' >> /etc/apk/repositories
RUN command echo 'http://dl-cdn.alpinelinux.org/alpine/v3.9/community' >> /etc/apk/repositories
RUN apk update && apk upgrade

# MongoDB - Database
RUN apk add openrc mongodb mongodb-tools
RUN command mkdir -p /data/db/ && chown root /data/db
CMD "rc-update add mongodb default && rc-service mongodb start -q"

# ClamAV - Anti Virus
RUN apk add clamav clamav-daemon
RUN command freshclam
RUN command mkdir "/home/log"
RUN command mkdir "/home/scan"

# Micronaut - Rest Api
RUN apk add openjdk17
COPY ./build/docker/main/layers/ /home/app/

# Shell - Start script
COPY init.sh usr/bin/init.sh
ENTRYPOINT ["sh", "./usr/bin/init.sh"]


