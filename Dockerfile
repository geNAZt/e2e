FROM ubuntu:bionic

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN dpkg --add-architecture i386 && export DEBIAN_FRONTEND=noninteractive DEBCONF_NONINTERACTIVE_SEEN=true && \
    apt-get update && apt-get install -y xvfb wget git gcc software-properties-common build-essential x11proto-core-dev libx11-dev \
    libegl1-mesa:i386 libgl1-mesa-glx:i386 libxtst-dev libxkbcommon-x11-dev libx11-xcb-dev libpng-dev xfce4 firefox x11vnc tesseract-ocr && \
    wget -O - https://mcpelauncher.mrarm.io/apt/conf/public.gpg.key | apt-key add - && \
    add-apt-repository 'deb http://mcpelauncher.mrarm.io/apt/ubuntu/ bionic main' && \
    apt-get install -y msa-daemon msa-ui-qt mcpelauncher-client mcpelauncher-ui-qt && \
    wget -P /tmp https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf /tmp/go1.12.4.linux-amd64.tar.gz && \
    rm -rf /tmp/* && mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH" && \
    rm -rf /var/lib/apt/lists/*

ADD entrypoint.sh /root/entrypoint.sh

VOLUME ["/go/src"]