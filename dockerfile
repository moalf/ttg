FROM golang:1.17-alpine

WORKDIR /ttg

COPY . /ttg

USER root

RUN apk add bash --no-cache bash

CMD ["/bin/sh" "-c" "tail -f /dev/null"]