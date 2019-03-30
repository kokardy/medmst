from alpine:3.9

RUN apk --update add wget grep bash httpie supervisor && rm -rf /var/cache/apk/*
RUN mkdir -p /asset/data
COPY bin /asset/bin
COPY http /http
COPY daily /var/spool/cron/crontabs/daily
COPY supervisord.conf /etc/supervisord.conf
RUN cd /asset/bin && ./Init.sh /asset/data
CMD supervisord -c /etc/supervisord.conf
