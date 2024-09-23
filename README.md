# Stampy

[![Test](https://github.com/Jelloeater/stampy/actions/workflows/test.yml/badge.svg)](https://github.com/Jelloeater/stampy/actions/workflows/test.yml)
![coverage](https://raw.githubusercontent.com/Jelloeater/stampy/refs/heads/badges/.badges/main/coverage.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/jelloeater/stampy)](https://goreportcard.com/report/github.com/jelloeater/stampy)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/jelloeater/stampy)

![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/jelloeater/stampy/total)
![GitHub Release](https://img.shields.io/github/v/release/jelloeater/stampy)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jelloeater/stampy)
![GitHub Release Date](https://img.shields.io/github/release-date/jelloeater/stampy)

**<https://libraries.io/go/github.com%2FJelloeater%2Fstampy>**

![Stampy](stampy.gif)

## Install

### Apt (Perfered)

https://packagecloud.io/jelloeater/stampy

```shell
curl -s https://packagecloud.io/install/repositories/jelloeater/stampy/script.deb.sh | sudo bash
sudo apt-get install stampy
```

### Yum (Perfered)

```shell
curl -s https://packagecloud.io/install/repositories/jelloeater/stampy/script.rpm.sh | sudo bash
sudo yum install stampy
```

### Binary (eget)

Use <https://github.com/zyedidia/eget>

```shell
curl https://zyedidia.github.io/eget.sh | sh
sudo mv eget /usr/local/bin
sudo eget jelloeater/stampy --to /usr/local/bin --asset=tar.gz
```

### Via Go

```shell
go install github.com/Jelloeater/stampy@latest
```

## Usage

### Settings

You can use ENV_VARs to override settings

```shell
export STAMPY_TZ='EST'
export STAMPY_FORMAT='01/02/2006 15:04'
export STAMPY_NTP='pool.ntp.org'
```

### Help

```shell
❯ stampy -h
NAME:
   stampy - Copy formatted timestamp to system clipboard

USAGE:
   stampy [global options] command [command options] [arguments...]

VERSION:
   v1.1.0 build 3affdf7d46a8c82a0e0f34be2f6af9d8a10f7eac

AUTHOR:
   Jelloeater

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value      Timestamp format (default: "2006-01-02T15:04:05Z07:00")
   --timezone value    Timezone (default: "UTC")
   --ntp_server value  NTP Server (ex pool.ntp.org)
   --help, -h          show help
   --version, -v       print the version
```

## Build

Run

```shell
go build -o ./build .
./build/stampy
```

or use <https://taskfile.dev/>

## Previous Work

<https://github.com/oz/tz>
