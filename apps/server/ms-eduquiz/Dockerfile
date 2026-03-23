FROM openjdk:17-jdk-slim-buster
WORKDIR /home/app
COPY ./build/docker/main/layers/libs /home/app/libs
COPY ./build/classes /home/app/classes
COPY ./build/docker/main/layers/resources /home/app/resources
COPY ./build/docker/main/layers/app/application.jar /home/app/application.jar
EXPOSE 8086
ENTRYPOINT ["java", "-jar", "/home/app/application.jar"]
