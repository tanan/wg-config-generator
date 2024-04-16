FROM ubuntu:24.04 as builder

ARG S6_OVERLAY_VERSION=3.1.6.2

RUN apt-get update && apt-get install -y --no-install-recommends \
    wireguard \
    iproute2 \
    iptables \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

EXPOSE 51820/udp
ENTRYPOINT ["/init"]
