FROM golang:1.23.0-bookworm

WORKDIR /root

RUN apt update && apt install -y git unzip
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip" && \
  unzip awscliv2.zip && \
  ./aws/install

WORKDIR /go/src/github.com/yt-d-rep/event-sourcing-go