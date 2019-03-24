from alpine:3.9

RUN apk --update add wget grep bash && rm -rf /var/cache/apk/*
RUN mkdir -p /asset
COPY bin /asset/bin
RUN cd /asset/bin && ./Init.sh
