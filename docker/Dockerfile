FROM alpine:latest
RUN mkdir /thingmodels
WORKDIR /thingmodels
COPY tm_catalog_cli /thingmodels
RUN chmod 755 /thingmodels/tm_catalog_cli
ENTRYPOINT ["/thingmodels/tm_catalog_cli", "serve",  "--host", "0.0.0.0", "--port", "8080"]
