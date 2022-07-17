# syntax=docker/dockerfile:1
FROM golang:1.16.7-alpine3.14 as builder
ENV GOOS=linux
RUN apk add --update git openssh-client && rm -rf /var/cache/apk/* && \
    mkdir /root/.ssh && echo "StrictHostKeyChecking no" > /root/.ssh/config && \
    echo "${SSH_KEY}" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    git clone https://github.com/gopalrg310/json-diff.git /json-diff
WORKDIR /json-diff
COPY . .
RUN go mod download && \
    go build -o json-diff main.go
