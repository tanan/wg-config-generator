# wg-config-generator

This is a tool for generating Wireguard configuration automatically.

## build

```
go build main.go -o wgconf
```

## dependencies

- Wireguard

## preparation

- create a `config.yaml` file like [this](./config.yaml).
- create public and private keys for vpn server and set publickey string and privatekey file path to your `config.yaml`.

```
wg genkey | tee privatekey | wg pubkey > publickey
chmod 600 privatekey
```

## usage

- create a client profile

```
./wgconf client create [user_name] -i [ip_address] -t [output_type] -c [config_path]
```

- create a server config

```
./wgconf client server create -c [config_path]
```
