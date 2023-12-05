FROM alpine:latest
RUN mkdir /thingmodels
WORKDIR /thingmodels
COPY tm-catalog-cli /thingmodels
RUN chmod 755 /thingmodels/tm-catalog-cli
ENTRYPOINT ["/thingmodels/tm-catalog-cli", "serve",  "--host", "0.0.0.0", "--port", "8080"]
