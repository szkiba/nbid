# nbid

[![codecov](https://codecov.io/gh/szkiba/nbid/branch/master/graph/badge.svg)](https://codecov.io/gh/szkiba/nbid)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/szkiba/nbid)
[![Go Report Card](https://goreportcard.com/badge/github.com/szkiba/nbid)](https://goreportcard.com/report/github.com/szkiba/nbid)

**Name Based globally unique ID generator**

**NBID** can be generate from arbitrary names using cryptographic hash function (SHA256 by default). The binary representation of the NBID is a 16 byte byte array. The string representation is using base32 hex (w/o padding) for space efficiency when stored in that form (26 bytes). The hex variant of base32 is used to retain the sortability (string and binary representation has same order).

NBID doesn't use base64 because case sensitivity and the 2 non alphanum chars may be an issue when transported as a string between various systems. To validate a base32 NBID, expect a 26 chars long, all uppercase sequence of `A` to `V` letters and `0` to `9` numbers (`[0-9A-V]{26}`).

## Features

- *Easy to use* - Simly a 128 bit hash without complexity of UUID.
- *Modern hash* - You are no limited to obsoloted MD5/SHA1 hash usd by UUID. Defaul hash function is SHA256, but you can use any hash with at least 128 bit output.
- *Compact* - String representation size is only 26 characters instead of UUID's 35 characters.
- *Sortable* - Binary and string representation has same order (because of base32hex)
- *No case sensitivity* - Only uppercase letters and numbers make it usable environment without real case sensitive names (DNS host names, FAT filesystem, etc).

## Install

> This section about CLI tool, for API usage check [Documentation](https://pkg.go.dev/github.com/szkiba/nbid).

You can install the pre-compiled binary or use Docker.

### Install the pre-compiled binary

Download the pre-compiled binaries from the [releases page](https://github.com/szkiba/nbid/releases) and
copy to the desired location.

### Install with Go

If you have Go environment set up, you can build nbid from source by running:

```sh
go get github.com/szkiba/nbid/cmd/nbid
```

Binary would be installed to $GOPATH/bin/nbid.

### Running with Docker

You can also use it within a Docker container. To do that, you'll need to
execute something more-or-less like the following:

```sh
docker run --network=host szkiba/nbid
```

### Verifying your installation

To verify your installation, use the `nbid -v` command:

```sh
$ nbid -v

nbid/1.0.0 linux/amd64
```

You should see `nbid/VERSION` in the output.

### Usage

To print usage information, use the `nbid --help` command:

```
$ nbid --help

usage: ./nbid [name]

Generate NBID for name, or random NBID if name is missing.

Example: ./nbid "The quick brown fox jumps over the lazy dog"
Output: QUKFNCO7QU098QEAJAUB021E9S

  -v    prints version
```

## TODO

Document, document, document...
