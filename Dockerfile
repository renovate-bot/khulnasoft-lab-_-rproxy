FROM golang:1.22.1-alpine3.19 AS base

RUN apk --update upgrade && apk --no-cache --update-cache --upgrade --latest add ca-certificates build-base gcc

WORKDIR /build

ADD go.mod go.mod
ADD go.sum go.sum

ENV GO111MODULE on
ENV CGO_ENABLED 1

RUN go mod download

ADD . .

ARG VERSION

RUN go build  \
    -ldflags="-X main.version=${VERSION}" \
    -o /usr/bin/rproxy

FROM alpine:3.16

RUN addgroup -S rproxy; \
    adduser -S rproxy -G rproxy -D -u 10000 -s /bin/nologin;

COPY --from=base /usr/bin/rproxy /usr/bin/rproxy

USER 10000

ENTRYPOINT ["rproxy"]
CMD ["start", "--config", "/etc/rproxy/config.yaml"]
