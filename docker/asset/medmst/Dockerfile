FROM ubuntu:16.04

COPY asset /asset

RUN apt-get update && \
    apt-get install -y golang git python jlha && \
    go install github.com/kokardy/medmst && \
    mkdir /bootstrap && \
    /asset/routine.sh 

