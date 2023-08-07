# Marketplace infrastructure

The marketplace infrastructure has been designed to implement the microservices archquitecture.

## Infrastructure map

![Infrastructure map](https://drive.google.com/uc?export=view&id=1cd8k0jgeQ91mPj5apI2aWwrTDqlzVgu1)

The idea is simple, we must use an gateway service to redirect the users requests on each API that it need.

The gateway is a nginx proxy service either in docker context or kubernetes context (ingress controller nginx).

Each API must connect to a mongo database to perisist the data.
