FROM ubuntu:20.04

WORKDIR /

RUN apt-get update && \
    apt-get install curl --yes && \
    curl -LO https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    apt-get install vim --yes

COPY config ./root/.kube/config

