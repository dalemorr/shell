# shell

This is a program that can perform operations on a Flinstone disk.

<details>
    <summary>Table of Contents</summary>

- [shell](#shell)
  - [Installation](#installation)
  - [Usage](#usage)
</details>

## Installation

To install this program, run:

```sh
go get -u github.com/dalemorr/shell
go install github.com/dalemorr/shell
```

To ensure that installation was successful, run:

```sh
shell -v
```

This should print the current version number. If instead an error message is displayed saying `Command 'shell' not found`, you may need to add GOROOT and GOPATH to your PATH. To do this on Linux, append the following to your `.bashrc`:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## Usage

To see usage details, run:

```sh
shell -h
```

It will display the following message:

```sh

```