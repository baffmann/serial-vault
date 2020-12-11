FROM golang

#ENV http_proxy http://
#ENV https_proxy http:/
ENV CSRF_SECURE=disable

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y postgresql-client ca-certificates

ADD . /go/src/github.com/CanonicalLtd/serial-vault

WORKDIR /go/src/github.com/CanonicalLtd/serial-vault
RUN go get ./...

COPY settings.yaml /go/src/github.com/CanonicalLtd/serial-vault
COPY ./docker-compose/docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
