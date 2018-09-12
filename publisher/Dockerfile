FROM alpine

RUN apk update && apk add ca-certificates

COPY ./bin/publisher /usr/bin/publisher

EXPOSE 8091

ENTRYPOINT ["publisher"]