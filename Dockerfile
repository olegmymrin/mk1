FROM alpine as base

WORKDIR ./data/api

COPY ./data/api/api-service ./api-service
COPY ./conf/config.yaml ./config.yaml

CMD ["./api-service", "./config.yaml"]