FROM ubuntu:20.04

RUN apt-get update

RUN apt-get -y --no-install-recommends install \
	nginx \
	golang

RUN apt-get clean

WORKDIR /

ENV GOPATH /go
ENV PATH ${PATH}:genomdb

RUN nginx -c /nginx/server.conf
