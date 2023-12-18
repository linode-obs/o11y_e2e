ARG ARCH="amd64"
ARG OS="linux"
FROM golang:alpine3.18 AS builderimage
LABEL maintainer="Akamai SRE Observability Team <support@linode.com>"
WORKDIR /go/src/go_project_template
COPY . .
RUN go build -o go_project_template cmd/main.go

###################################################################

FROM golang:alpine3.18
COPY --from=builderimage /go/src/go_project_template/go_project_template /app/
WORKDIR /app

EXPOSE      9141
USER        nobody
ENTRYPOINT  [ "./go_project_template" ]
