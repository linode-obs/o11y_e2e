This is an opinionated Golang project template skeleton. Replace linode-obs/go_project_template with your project. Sometimes just go_project_template will need to be replaced. Also search TODO.

Make sure to write tests and `pre-commit install`. You'll also want to install [goreleaser](https://goreleaser.com/) and [pre-commit](https://pre-commit.com/). Note that the MIT license is included too.

Beware that:

* The Go version is set at `1.21` in [golang-tests.yaml](.github/workflows/golang-tests.yaml). This could be automated with renovate or dependabot
* Markdown table of contents be nice to add, easy to create with an [extension](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one) in your editor of choice
* Document your CLI values and metrics instrumented

# go_project_template

![Github Release Downloads](https://img.shields.io/github/downloads/linode-obs/go_project_template/total.svg)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/linode-obs/go_project_template/blob/master/LICENSE)
[![golangci-lint](https://github.com/linode-obs/go_project_template/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/linode-obs/go_project_template/actions/workflows/golangci-lint.yaml)
![Go Report Card](https://goreportcard.com/badge/github.com/linode-obs/go_project_template)
[![contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat")](https://github.com/linode-obs/go_project_template/issues)

Project description.

We recommend using a standard [go project layout](https://github.com/golang-standards/project-layout) too.

Features:

* Feature 1
* Feature 2

## Installation

### Debian/RPM package

Substitute `{{ version }}` for your desired release.

```bash
wget https://github.com/linode-obs/go_project_template/releases/download/v{{ version }}/go_project_template_{{ version }}_linux_amd64.{deb,rpm}
{dpkg,rpm} -i go_project_template_{{ version }}_linux_amd64.{deb,rpm}
```

### Docker

```console
sudo docker run \
--privileged \
ghcr.io/linode-obs/go_project_template
```

### Binary

```bash
wget https://github.com/linode-obs/go_project_template/releases/download/v{{ version }}/go_project_template_{{ version }}_Linux_x86_64.tar.gz
tar xvf go_project_template_{{ version }}_Linux_x86_64.tar.gz
./go_project_template/go_project_template
```

### Source

```bash
wget https://github.com/linode-obs/go_project_template/archive/refs/tags/v{{ version }}.tar.gz
tar xvf go_project_template-{{ version }}.tar.gz
cd ./go_project_template-{{ version }}
go build go_project_template.go
./go_project_template.go
```

## Releasing

1. Merge commits to main.
2. Tag release `git tag -a v1.0.X -m "message"`
3. `git push origin v1.0.X`
4. `goreleaser release`

## Contributors

Contributions welcome! Make sure to `pre-commit install`
