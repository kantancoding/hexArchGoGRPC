FROM golang:1.15.3-alpine3.12

EXPOSE 8080

RUN apk update \
    && apk add --no-cache \
        mysql-client \
        build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY ./grpc_entrypoint.sh /usr/local/bin/grpc_entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/grpc_entrypoint.sh

RUN go build cmd/grpc/server.go
RUN mv server /usr/local/bin/

CMD ["server"]
ENTRYPOINT ["grpc_entrypoint.sh"]
