# DDNS Daemon ![GitHub](https://img.shields.io/github/license/nmapx/ddns-daemon?style=flat) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/nmapx/ddns-daemon?style=flat) [![Go Report Card](https://goreportcard.com/badge/github.com/nmapx/ddns-daemon)](https://goreportcard.com/report/github.com/nmapx/ddns-daemon)

OVH DynHost self-hosted daemon. Multiple hosts supported!

## How it works

It's fetching your IP address from https://ifconfig.co API then updating all your OVH 
DynHost hosts based on the configuration file.

## Config

By default `config.yml` is expected to be in the same directory but you can customize
it with `--config-filepath` param.

```yaml
hosts:
    first_host:
        host: first.host.dynhost.ovh
        user: firstHostUsername
        pass: firstHostPassword
    second_host:
        host: second.host.dynhost.ovh
        user: secondHostUsername
        pass: secondHostPassword
    ...
```

## Quick setup

### Production

1. Download [latest](/nmapx/ddns-daemon/releases/latest) executable
2. Prepare yaml config file
3. Configure the daemon (eg. service) or run it manually (eg. screen)

### Development

Run it with or without Docker.

1. Clone the repository
2. Create `.env` file based on `.env.dist`
3. Prepare yaml config file
4. Build docker image and run Golang environment
5. Contribute, build binary, run daemon

## License

[MIT License](./LICENSE)
