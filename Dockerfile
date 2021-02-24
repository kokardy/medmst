from golang:alpine

RUN apk --update add wget grep bash httpie supervisor git coreutils && rm -rf /var/cache/apk/*
RUN go get github.com/gin-gonic/gin
RUN mkdir -p /asset/data
COPY bin /asset/bin
COPY http /http
COPY daily /asset/
RUN cd /http && go build -o server .
RUN cat /asset/daily >> /var/spool/cron/crontabs/root
COPY supervisord.conf /etc/supervisord.conf
CMD supervisord -c /etc/supervisord.conf
