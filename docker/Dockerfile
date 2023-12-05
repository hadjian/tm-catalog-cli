FROM alpine:latest
RUN mkdir /thingmodels
WORKDIR /thingmodels
RUN /bin/sh download-cli.sh
ENTRYPOINT ["/thingmodels/tm-catalog-cli", "serve",  "--host", "0.0.0.0", "--port", "8080"]
