# o11y_e2e

![Github Release Downloads](https://img.shields.io/github/downloads/linode-obs/o11y_e2e/total.svg)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/linode-obs/o11y_e2e/blob/master/LICENSE)
[![golangci-lint](https://github.com/linode-obs/o11y_e2e/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/linode-obs/o11y_e2e/actions/workflows/golangci-lint.yaml)
![Go Report Card](https://goreportcard.com/badge/github.com/linode-obs/o11y_e2e)
[![contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat")](https://github.com/linode-obs/o11y_e2e/issues)

The best way to measure pipeline effectiveness, especially for SLOs, is examining behavior from the end user's perspective. This project aims to simplify that by ingesting a log into systems like Loki and determining how long it takes for that metric to be query-able as an end-user.

This project is inspired by Alex Hidalgo's *[Implementing Service Level Objectives](https://learning.oreilly.com/library/view/implementing-service-level/9781492076803/ch01.html#:-:text=you%20should%20consider%20measuring,reaches%20the%20other%20end)*:

>... even better would be to measure how much time elapses after inserting a particular record into the pipeline before you can retrieve that exact record at the other end

This project generates metrics on an internal interval rather than on-scrape as this type of data is not always going to be a good fit during Prometheus scrape periods (~`2m` is common).

- [o11y\_e2e](#o11y_e2e)
  - [Example Configuration](#example-configuration)
  - [Installation](#installation)
    - [Debian/RPM package](#debianrpm-package)
    - [Docker](#docker)
    - [Binary](#binary)
    - [Source](#source)
  - [Potential Features](#potential-features)
    - [Proposed features](#proposed-features)
    - [Vaguely planned features](#vaguely-planned-features)
    - [Not planned features](#not-planned-features)
  - [Releasing](#releasing)
  - [Contributors](#contributors)

## Example Configuration

```yaml
o11y_e2e:
  - name: loki_prod
      type: loki # only loki is currently supported
      enabled: true
      url: "https://my_loki_api_url.com"
      timeout: "1m" # how long to wait before giving up querying the metric
      tls_config: # tls.Config standard options
          insecure_skip_verify: false
      labels:
      - cluster: infra-logging-atl1-us-staging
        dc: atl1
        environment: production
      syslog_push: # rather than just generate log to stdout/journalctl, send logs to syslog reciever
          enabled: true # TODO - document defaults and structure better
          url: "my_syslog_reciever.com"
          tls_config:
            insecure_skip_verify: false
  - name: loki_prod_region2
    # etc..
```

These config options will result in a logfmt log line of:

```console
creation_time=11:11:11, level=info, application=o11y_e2e, hash=MYHASH1234AFFF, cluster=infra-logging-atl1-us-staging, dc=atl1, environment=production, syslog=true, syslog_url="my_syslog_reciever.com"
```

## Installation

### Debian/RPM package

Substitute `{{ version }}` for your desired release.

```bash
wget https://github.com/linode-obs/o11y_e2e/releases/download/v{{ version }}/o11y_e2e_{{ version }}_linux_amd64.{deb,rpm}
{dpkg,rpm} -i o11y_e2e_{{ version }}_linux_amd64.{deb,rpm}
```

### Docker

```console
sudo docker run \
--privileged \
ghcr.io/linode-obs/o11y_e2e
```

### Binary

```bash
wget https://github.com/linode-obs/o11y_e2e/releases/download/v{{ version }}/o11y_e2e_{{ version }}_Linux_x86_64.tar.gz
tar xvf o11y_e2e_{{ version }}_Linux_x86_64.tar.gz
./o11y_e2e/o11y_e2e
```

### Source

```bash
wget https://github.com/linode-obs/o11y_e2e/archive/refs/tags/v{{ version }}.tar.gz
tar xvf o11y_e2e-{{ version }}.tar.gz
cd ./o11y_e2e-{{ version }}
go build o11y_e2e.go
./o11y_e2e.go
```

## Potential Features

### Proposed features

- Loki support
- Generate a unique log and provide a slew of Prometheus metrics on the latency and durability of that log message at the end of the pipeline
- Support syslog-style "push" to a pipeline entry point rather than rely on unknown external sending feature like promtail or otelcol agent
- Support multiple pipelines from one instance of o11y_e2e

### Vaguely planned features

- Expansion into metrics pipelines like Thanos, Mimir, Victoria Metrics
- Support for elasticsearch or other log aggregators

### Not planned features

- Prometheus multi target exporter pattern as the data generation can be too slow and managing both internal metric generation and on-demand scrapes is painful

## Releasing

1. Merge commits to main.
2. Tag release `git tag -a v1.0.X -m "message"`
3. `git push origin v1.0.X`
4. `goreleaser release`

## Contributors

Contributions welcome! Make sure to `pre-commit install`
