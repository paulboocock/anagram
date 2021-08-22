# Anagram Finder

[![Build](https://github.com/paulboocock/anagram/actions/workflows/build.yml/badge.svg)](https://github.com/paulboocock/anagram/actions/workflows/build.yml)
[![Version (latest semver)](https://img.shields.io/github/v/release/paulboocock/anagram)](https://github.com/paulboocock/anagram/releases)
[![License](https://img.shields.io/github/license/paulboocock/anagram)](LICENSE)
![Go Version](https://img.shields.io/github/go-mod/go-version/paulboocock/anagram)

This application finds anagrams from a sorted list of strings and prints them to stdout.

## Quick Start

### Docker

Using [Docker](https://www.docker.com/get-started) you can run the application with a mounted `/data` directory (Linux/MacOS/Powershell):

```bash
docker run --mount type=bind,src=$(pwd)/data,dst=/data ghcr.io/paulboocock/anagram:latest data/example1.txt
```

### Binaries

A number of binaries are available to download, hosted via [GitHub Releases](https://github.com/paulboocock/anagram/releases):

| Linux (amd64) |  Linux (arm64) |  MacOS (Intel) | MacOS (Apple Silicon) | Windows (x64) |
|---|---|---|---|---|
| [Download][linux_amd64] | [Download][linux_arm64] | [Download][darwin_amd64] | [Download][darwin_arm64] | [Download][windows_amd64] |

Once downloaded, unzip the executable.

Linux/MacOS (update the executable name to match the downloaded version):

```bash
./anagram data/example1.txt
```

Windows:

```bash
anagram_windows_amd64.exe data/example1.txt
```

## Maintainer quick start

Requires `go` and `make` to be installed.

Go: [Download](https://golang.org/dl/)

`make` is available on Linux and MacOS.  
Windows users can install via [chocolatey](https://chocolatey.org/install) (`choco install make`) or use [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install-win10).

### Building

To build the cli application:

```bash
make cli
```

To build the docker container:

```bash
make container
```

### Testing

To run all tests:

```bash
make test
```

### Linting

To lint all files

```bash
make lint
```

[linux_amd64]: https://github.com/paulboocock/anagram/releases/download/1.0.0/anagram_linux_amd64.zip
[linux_arm64]: https://github.com/paulboocock/anagram/releases/download/1.0.0/anagram_linux_arm64.zip
[darwin_amd64]: https://github.com/paulboocock/anagram/releases/download/1.0.0/anagram_darwin_amd64.zip
[darwin_arm64]: https://github.com/paulboocock/anagram/releases/download/1.0.0/anagram_darwin_arm64.zip
[windows_amd64]: https://github.com/paulboocock/anagram/releases/download/1.0.0/anagram_windows_amd64.zip