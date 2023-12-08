FROM alpine:latest

ENV SERVER_IP=localhost
ENV SERVER_PORT=8080

RUN apk update \
 && apk add --no-cache curl jq \
 && rm -rf /var/cache/apk/*

RUN mkdir /thingmodels
WORKDIR /thingmodels
COPY ./docker/download-cli.sh /
RUN cd /; chmod 755 ./download-cli.sh; ./download-cli.sh
RUN mv /tm-catalog-cli /thingmodels/
RUN chmod 755 /thingmodels/tm-catalog-cli
RUN /thingmodels/tm-catalog-cli remote add --type file catalog /thingmodels
RUN /thingmodels/tm-catalog-cli remote set-default catalog
ENTRYPOINT /thingmodels/tm-catalog-cli serve --host ${SERVER_IP} --port ${SERVER_PORT}
