FROM ubuntu:16.04

RUN rm /etc/apt/sources.list && \
    echo 'deb http://mirrors.aliyun.com/ubuntu/ xenial main restricted' > /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y quagga && \
    apt-get clean

COPY ./* /etc/quagga/
