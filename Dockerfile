from alpine:3.9

RUN apk --update add wget grep bash httpie && rm -rf /var/cache/apk/*
RUN mkdir -p /asset/data
COPY bin /asset/bin
COPY daily /var/spool/cron/crontabs/daily
RUN cd /asset/bin && ./Init.sh /asset/data
CMD crond -l 2 -f
