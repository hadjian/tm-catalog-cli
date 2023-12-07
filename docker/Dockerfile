FROM alpine:latest
RUN mkdir /thingmodels
WORKDIR /thingmodels
COPY tm-catalog-cli /thingmodels
RUN chmod 755 /thingmodels/tm-catalog-cli
RUN /thingmodels/tm-catalog-cli remote add --type file catalog /thingmodels
RUN /thingmodels/tm-catalog-cli remote set-default catalog
ENTRYPOINT ["/thingmodels/tm-catalog-cli", "serve",  "--host", "0.0.0.0", "--port", "8080"]
