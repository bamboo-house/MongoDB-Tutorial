FROM golang:1.19

RUN apt-get update && \
  apt-get install -y \
  git \
  curl

# clean
RUN apt-get clean && \
  rm -rf /var/lib/apt/lists/*

RUN mkdir -p /go/src/app
COPY ./app /go/src/app

WORKDIR /go/src