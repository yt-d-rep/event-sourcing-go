FROM golang:1.23.0-bookworm

RUN apt update && apt install -y git