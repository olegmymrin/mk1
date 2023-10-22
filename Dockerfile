FROM alpine as base

WORKDIR ./data/api

COPY ./data/api/api-service ./api-service

CMD ["./api-service"]