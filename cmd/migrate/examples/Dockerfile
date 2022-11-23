FROM ubuntu:bionic

RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent

RUN curl -sSL https://packagecloud.io/morigs/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/morigs/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && \
    apt-get install -y migrate

RUN migrate -version
