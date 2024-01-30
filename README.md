# SZ
Prints the size of a file in human readable format.

## Why?
Because I am on windows and it is stupid


## Installation
```console
$ go get github.com/elliot40404/sz
```

```console
$ scoop install https://raw.githubusercontent.com/elliot40404/sz/master/sz.json
```

```console
$ scoop bucket add elliot https://github.com/elliot40404/elliot
$ scoop install elliot/sz
```

```console
$ git clone
$ cd sz
$ go install cmd/sz or make sz-install
# if you have tinygo installed you get a smaller binary and faster binary
$ make sz-tiny
$ cp bin/sz.exe <your GOPATH or preffered binary location>
```

## Usage

```console
$ sz
sz version v0.2.0
Prints file sizes in bytes, kilobytes, megabytes and gigabytes
Author: Elliot40404
Usage: sz <file> <file> <file>
```

```console
$ sz main.go
715 B  0.70 KB  0.001 MB  0.000 GB
```

```console
$ sz README.md sz.json
724 B  0.71 KB  0.001 MB  0.000 GB README.md
361 B  0.35 KB  0.000 MB  0.000 GB sz.json
```

## License
MIT