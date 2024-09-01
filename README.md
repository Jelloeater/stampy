# Stampy

## Usage

```shell
stampy --help
NAME:
   stampy - Copy formatted timestamp to system clipboard

USAGE:
   stampy [global options] command [command options] [arguments...]

VERSION:
   vdev build none

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value    Timestamp format (default: "2006-01-02T15:04:05Z07:00")
   --timezone value  Timezone (default: "UTC")
   --help, -h        show help
   --version, -v     print the version
```

## Build

Run

```shell
go build -o ./build .
./build/stampy
```

or use https://taskfile.dev/

