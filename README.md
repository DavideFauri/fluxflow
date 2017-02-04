# FluxFlow

A command line tool to adjust [f.lux](https://justgetflux.com/) settings (on OS X).

## Installation

Get the binary for OS X from the latest [release].

Or use `go get`:

```
go get github.com/tonglil/fluxflow
```

[release]: https://github.com/tonglil/fluxflow/releases

## Usage

While f.lux is running:

```bash
$ fluxflow
A command line tool to adjust f.lux settings

Usage:
  fluxflow [command]

Available Commands:
  completion  Output shell completion code for tab completion
  darkroom    Set f.lux to darkroom mode
  hour        Disable f.lux for 1 hour
  movie       Set f.lux to movie mode for 2.5 hours
  sunrise     Disable f.lux until sunrise

Use "fluxflow [command] --help" for more information about a command.
```

## Tab completion

```bash
$ source <(fluxflow completion)
```
