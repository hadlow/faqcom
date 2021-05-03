FROM ubuntu:20.04

RUN apt-get update

RUN apt-get -y --no-install-recommends install \
	nginx \
	golang

RUN apt-get clean

WORKDIR /

ENV GOPATH /go
ENV PATH ${PATH}:genomdb

COPY src /genomdb

COPY /server/goweb.service /lib/systemd/system

RUN service goweb start
