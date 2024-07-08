ARG VERSION=1.19
FROM golang:${VERSION}-alpine

RUN set -ex; \
    apk update && apk add --no-cache \
        git 

WORKDIR /workspace
COPY . .