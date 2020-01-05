FROM golang:1.13.5-alpine3.11

RUN apk add make cmake

WORKDIR /usr/local/src

COPY . /usr/local/src

RUN make build

RUN mv /usr/local/src/jsonlinter/linter /usr/local/bin/linter

ENTRYPOINT [ "linter" ]