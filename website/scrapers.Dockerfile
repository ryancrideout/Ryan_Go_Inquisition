FROM ubuntu:25.10

RUN apt update && apt -y upgrade && apt install -y curl locales && locale-gen en_US.UTF-8

# Install necessary libraries
RUN apt install -y \
    wget \
    # Chrome dependencies
    libappindicator3-1 \
    libgtk-3-0 \
    libxss1 \
    libgbm1 \
    libnss3 \
    libasound2t64 \
    libnspr4 \
    fonts-liberation \
    libu2f-udev \
    libvulkan1 \
    xdg-utils 

# Install Go and configure some of it's settings.
RUN curl -o go.tar.gz https://dl.google.com/go/go1.24.3.linux-amd64.tar.gz && tar -xf go.tar.gz
ENV PATH=$PATH:/go/bin

# Install Google Chrome
RUN wget --quiet https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -O google-chrome-stable.deb && \
    dpkg -i google-chrome-stable.deb 

WORKDIR /opt/scrapers/web_scrapers