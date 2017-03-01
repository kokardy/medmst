FROM ubuntu:16.04

COPY asset /asset

RUN apt-get update && \
    apt-get install -y golang python jlha-utils && \
    apt-get clean

ENV GOPATH=/go
RUN mkdir /go && mkdir /bootstrap
RUN go install github.com/kokardy/medmst && \
    sh /asset/routine.sh 

