# DDNS Daemon

OVH DynHost self-hosted daemon. Multiple hosts supported!

## How it works

It's fetching your IP address from https://ifconfig.co API then updating all your OVH 
DynHost hosts based on the configuration file.

## Config

By default `config.yml` is expected to be in the same directory but you can customize
it with `--config` parameter.

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

1. Download latest binary
2. Prepare config file
3. Configure the daemon (service) or run it manually (screen)

### Development

Run it with or without Docker.

1. Clone the repository
2. Create .env file based on .env.dist
3. Prepare test config file
4. Build docker image and run Golang environment
5. Contribute, build binary, run daemon

## License

MIT
