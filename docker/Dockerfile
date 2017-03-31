FROM docker:1.13-dind



COPY asset /asset

WORKDIR /asset

#docker-compose
#ENV COMPOSE_URL=https://github.com/docker/compose/releases/download/1.11.2/docker-compose-Linux-x86_64
#RUN curl -L $COMPOSE_URL -o /usr/local/bin/docker-compose && \
#    chmod +x /usr/local/bin/docker-compose

RUN apk update && \
    apk add python python-dev py-pip build-base && \
    pip install docker-compose && \
    apk del python-dev build-base && \
    rm -r /var/cache/apk/*


ENTRYPOINT dockerd-entrypoint.sh
CMD ./compose.sh
