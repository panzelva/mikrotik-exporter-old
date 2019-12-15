## Prometheus Exporter for Mikrotik devices

**Important! This repository is based on https://github.com/nshttpd/mikrotik-exporter**

### Description

A Prometheus Exporter for Mikrotik devices that can be configured to collect metrics
from a single device or multiple devices. Single device monitoring can be configured
all on the command line. Multiple devices require a configuration file. A user will
be required that has read-only access to the device configuration via the API.

Currently the exporter collects metrics for interfaces and system resources. Others
can be added as long as published via the API.

### Mikrotik Config

Create a user on the device that has API and read-only access.

`/user group add name=prometheus policy=api,read,winbox`

Create the user to access the API via.

`/user add name=prometheus group=prometheus password=changeme`

### Single Device

`./mikrotik-exporter -address 10.10.0.1 -device my_router -password changeme -user prometheus`

where `address` is the address of your router. `device` is the label name for the device
in the metrics output to prometheus. The `user` and `password` are the ones you
created for the exporter to use to access the API.

### Config File

`./mikrotik-exporter -config-file config.yml`

where `config-file` is the path to a config file in YAML format.

#### example config

```yaml
devices:
  - name: my_router
    address: 10.10.0.1
    user: prometheus
    password: changeme
  - name: my_second_router
    address: 10.10.0.2
    user: prometheus2
    password: password_to_second_router

features:
  bgp: true
  dhcp: true
  dhcpv6: true
  dhcpl: true
  routes: true
  pools: true
  optics: true
```
