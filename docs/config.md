# kestrel Config Files

## Overview

kestrel config files are separated into two groups. 

/etc/kestrel/kestrel.conf (TOML format)
/etc/kestrel/peers.json

By separating the files, we can continue to support and interoperate with users of cjdns via the standard json-formatted peer credentials, but move new functionality into the human-readable TOML formatted main config file.

## kestrel.conf

This file is [TOML](https://github.com/toml-lang/toml)-formatted.

It has the following sections:

* Server
* Auth
* UDPInterface
* Admin
* Logging 

### Server section

The server section has the following minimum sections
* public_key
* private_key
* ipv6_address
* daemonize

Example:

    [server]
    public_key = "..."
    private_key = "..."
    ipv6_address = "..."
    daemonize = true

### Auth section

This section contains passwords for remote peers.

Example:

    [auth.peer1]
    password = "something very secure"

Other information can be included optionally:

    [auth.peer2]
    password = "something very secure"
    email = "peer2@example.com"


### Admin section

This section manages the admin RPC interface for kestrel.

TBC

### Logging section

This section manages log-related configuration

Example:

    [logging]
    enable = true
    level = info
    logger = file # or stderr
    path = /var/log/kestrel.log

## peers.json

TBC