ARG ARCH="amd64"
ARG OS="linux"
FROM golang:alpine3.18 AS builderimage
LABEL maintainer="Akamai SRE Observability Team <support@linode.com>"
WORKDIR /go/src/o11y_e2e
COPY . .
RUN go build -o o11y_e2e cmd/main.go

###################################################################

FROM golang:alpine3.18
COPY --from=builderimage /go/src/o11y_e2e/o11y_e2e /app/
WORKDIR /app

EXPOSE      9927
USER        nobody
ENTRYPOINT  [ "./o11y_e2e" ]
