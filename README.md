# vmagent-remotewrite-flag-renderer

## Description

Command line flag generator for [vmagent](https://docs.victoriametrics.com/vmagent/) remote write settings.

## Background

remote_write settings of Prometheu can be written in yaml, but vmagent settings must be specified as command line flags.

I needed a small command line tool to make migration easier.

See also

- https://docs.victoriametrics.com/vmagent/#advanced-usage 
- https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_write

## Install

```
go install github.com/johejo/vmagent-remotewrite-flag-renderer@latest
```

## Usage

```bash
vmagent $(vmagent-remotewrite-flag-renderer -config /path/to/vmagent.yaml -config-format=<prometheus or vmagent>)
```

### Prometheus Compatible File Format

See https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_write

Not all settings are supported. Settings that vmagent does not support are ignored.

```yaml
# snake_case
remote_write:
  - url: http://localhost:9000
    basic_auth:
      username: foo
      password: bar
  - url: http://localhost:9001
    basic_auth:
      username: hello
      password: world
```

will be converted to

```
-remoteWrite.url='http://localhost:9000' -remoteWrite.basicAuth.username='foo' -remoteWrite.basicAuth.password='bar' -remoteWrite.sendTimmeout='30s' -remoteWrite.url='http://localhost:9001' -remoteWrite.basicAuth.username='hello' -remoteWrite.basicAuth.password='world' -remoteWrite.sendTimmeout='30s'
```

Then, you can use with `vmagent` command like this.

```bash
vmagent $(vmagent-remotewrite-flag-renderer -config /path/to/prometheus.yaml -config-format=prometheus)
```

### VMAgent File Format

```yaml
# camelCase
remoteWrite:
  - url: http://localhost:9000
    basicAuth:
      username: foo
      password: hello
  - url: http://localhost:9001
    basicAuth:
      username: bar
      password: world
```

will be converted to

```
-remoteWrite.url='http://localhost:9000' -remoteWrite.basicAuth.username='foo' -remoteWrite.basicAuth.password='bar' -remoteWrite.sendTimmeout='30s' -remoteWrite.url='http://localhost:9001' -remoteWrite.basicAuth.username='hello' -remoteWrite.basicAuth.password='world' -remoteWrite.sendTimmeout='30s'
```

You can see the all settings by jsonschema with `vmagent-remotewrite-flag-renderer -vmagent-schema`.
