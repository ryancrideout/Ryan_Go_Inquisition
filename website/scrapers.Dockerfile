FROM ubuntu:25.10

RUN apt update && apt -y upgrade && apt install -y curl locales && locale-gen en_US.UTF-8

# Install Go and configure some of it's settings.
RUN curl -o go.tar.gz https://dl.google.com/go/go1.24.3.linux-amd64.tar.gz && tar -xf go.tar.gz
ENV PATH=$PATH:/go/bin

WORKDIR /opt/scrapers/web_scrapers