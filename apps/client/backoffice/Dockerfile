FROM nginx:1.24-alpine

RUN apk add bash

COPY card-pipeline.txt /etc/profile.d
COPY card-pipeline.sh /etc/profile.d

COPY /dist/bo /usr/share/nginx/html
RUN apk add --no-cache tzdata
ENV TZ=Europe/Lisbon


EXPOSE 80
