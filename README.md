# FluxFlow [![Build Status](https://travis-ci.org/DavideFauri/fluxflow.svg?branch=master)](https://travis-ci.org/DavideFauri/fluxflow)

A command line tool to adjust [f.lux](https://justgetflux.com/) settings (on OS X).
Original developer: [tonglil](github.com/tonglil)

## Installation

Get the binary for OS X from the latest [release].

Or use `go get`:

```
go get -u github.com/DavideFauri/fluxflow
```

[release]: https://github.com/DavideFauri/fluxflow/releases

## Usage

While f.lux is running:

```
$ fluxflow
A command line tool to adjust f.lux settings

Usage:
  fluxflow [command]

Available Commands:
	enable			Enable f.lux if disabled
  completion  Output shell completion code for tab completion
  darkroom    Set f.lux to darkroom mode
  hour        Disable f.lux for 1 hour
  movie       Set f.lux to movie mode for 2.5 hours
  sunrise     Disable f.lux until sunrise
  version     Print the version information

Use "fluxflow [command] --help" for more information about a command.
```

## Tab completion

```bash
source <(fluxflow completion)
```
