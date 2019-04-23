from golang:alpine

RUN apk --update add wget grep bash httpie supervisor git && rm -rf /var/cache/apk/*
RUN go get github.com/gin-gonic/gin
RUN mkdir -p /asset/data
COPY bin /asset/bin
COPY http /http
RUN cd /http && go build -o server .
RUN cat daily >> /var/spool/cron/crontabs/root
COPY supervisord.conf /etc/supervisord.conf
RUN cd /asset/bin && ./Init.sh /asset/data
CMD supervisord -c /etc/supervisord.conf
