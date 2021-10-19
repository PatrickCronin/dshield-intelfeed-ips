# dshield-intelfeel-ips

`dshield-intelfeel-ips` - print all IPs from the DShield API's Intelfeed to STDOUT

![Build](https://github.com/PatrickCronin/mmdb-dump-networks/workflows/Build/badge.svg)
![golangci-lint](https://github.com/PatrickCronin/mmdb-dump-networks/workflows/golangci-lint/badge.svg)
[![go report](https://goreportcard.com/badge/github.com/PatrickCronin/mmdb-dump-networks)](https://goreportcard.com/badge/github.com/PatrickCronin/dshield-intelfeel-ips)

* [Project Description](#project-description)
* [Usage](#usage)
* [Description](#description)
* [Installation](#installation)
* [Reporting Bugs and Issues](#reporting-bugs-and-issues)
* [Copyright and License](#copyright-and-license)

# Project Description

`dshield-intelfeel-ips` is a quick-and-dirty command line tool to retrieve and output the IPs currently on the [DShield API's Intelfeed](https://isc.sans.edu/api/#intelfeed).

# Usage

## Print Usage

```bash
$ ./dshield-intelfeed-ips -h
./dshield-intelfeed-ips
    Retrieve and output the IPs on DShield API's Intelfeed endpoint
    (https://isc.sans.edu/api/#intelfeed).

Usage: (-h|-email xxx@yyy.com)
$
```

## Print the IPs on the DShield API's Intelfeed

```bash
$ ./dshield-intelfeed-ips -email your@email.com
192.2.0.3
192.2.0.29
198.18.74.3
198.18.74.5
198.18.127.192
...
```

# Installation

## Binary Releases

Precompiled releases are currently available on our [Releases
page](https://github.com/PatrickCronin/dshield-intelfeed-ips/releases) for the
following platforms and architectures:

* Linux (i386 and x86_64)
* macOS (x86_64 and arm64)
* Windows (i386 and x86_64)

Look for a release that ends in .tar.gz or .zip. Download the release archive
for your platform and architecture.  Uncompress the archive and you'll see an
eponymous folder. In that folder, you'll find the `dshield-intelfeed-ips` program.
Copy that program to wherever you want it to live, and start using it.

## Linux Packages

Prebuilt packages are currently available on our [Releases
page](https://github.com/PatrickCronin/dshield-intelfeed-ips/releases) in the
following formats:

* .deb (Ubuntu or Debian)
* .rpm (RedHat or CentOS)

On Ubuntu or Debian, use `dpkg -i /path/to/the.deb` as root. On RedHat or
CentOS, `rpm -i /path/to/the.rpm` as root. `dshield-intelfeed-ips` will be
installed in to `/usr/bin/dshield-intelfeed-ips`.

## Building From Source

`dshield-intelfeed-ips` is written in Go, so you'll need a reasonably recent
version of Go (1.16+). This project aims to maintain support for the two most
recent major versions of the Go compiler.

With this in place, simply run:

```bash
$ go install github.com/PatrickCronin/dshield-intelfeed-ips/cmd/dshield-intelfeed-ips@latest
```

which will install `dshield-intelfeed-ips` into the directory named by the
`GOBIN` environment variable, which defaults to `$GOPATH/bin` or `$HOME/go/bin`
if the `GOPATH` environment variable is not set.

# Reporting Bugs and Issues

Bugs and other issues can be reported by filing an issue on our [GitHub issue
tracker](https://github.com/PatrickCronin/dshield-intelfeed-ips/issues).

# Copyright and License

This software is Copyright (c) 2021 by Patrick Cronin.

This is free software, licensed under the terms of the [MIT
License](https://github.com/PatrickCronin/dshield-intelfeed-ips/LICENSE.md).
