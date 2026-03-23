FROM registry.infra.atlanse.ddns.net/atl/images/openjdk:17-alpine

WORKDIR /home/app
COPY ./build/docker/main/layers/libs /home/app/libs
COPY ./build/classes /home/app/classes
COPY ./build/docker/main/layers/resources /home/app/resources
COPY ./build/docker/main/layers/app/application.jar /home/app/application.jar

COPY card-pipeline.txt /etc/profile.d
COPY card-pipeline.sh /etc/profile.d

RUN echo "/etc/profile.d/card-pipeline.sh" >> /root/.bashrc

ENV TZ=Europe/Lisbon
RUN rm -rf /var/cache/apt/*

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/home/app/application.jar"]
