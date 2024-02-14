# shell

This is a program that will read data a Flintstone disk and print its contents.

<details>
    <summary>Table of Contents</summary>

- [Build](#build)
- [Usage](#usage)
</details>

## Build

Go version 1.21.6 or later is expected for building this program. To build it, clone this repository to a local machine and run `go build .`

## Usage

To run this program, run `./shell <file-name>` to print data from a file named "file-name" or simply `./shell` to print data from Stdin. `disk.txt` is provided as an example file. that may be passed in. Run `./shell -v` for version information or `./shell -h` for a help message.