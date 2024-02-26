package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	AppName      = "shell"
	Version      = "v0.2.1"
	helpUsage    = "display help message"
	versionUsage = "display current version"
	printUsage   = "print raw contents of disk"
	dirUsage     = "list files on the disk"
	fileUsage    = "specify name of disk"
	diskHeader   = "XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	numRows    = 32
	numColumns = 64
	numBytes   = numRows * numColumns / 2
)

func main() {
	var err error

	// Handle args
	var isHelp bool
	var isVersion bool
	var isPrint bool
	var isDir bool
	var fileName string

	flag.BoolVar(&isHelp, "h", false, helpUsage)
	flag.BoolVar(&isHelp, "H", false, helpUsage)
	flag.BoolVar(&isHelp, "?", false, helpUsage)
	flag.BoolVar(&isVersion, "v", false, versionUsage)
	flag.BoolVar(&isVersion, "V", false, versionUsage)
	flag.BoolVar(&isPrint, "print", false, printUsage)
	flag.BoolVar(&isDir, "dir", false, dirUsage)
	flag.StringVar(&fileName, "f", "", fileUsage)
	flag.Parse()

	if isVersion {
		fmt.Printf("%s %s\n", AppName, Version)
		return
	}
	if isHelp {
		flag.CommandLine.Usage()
		return
	}

	// Read data
	var disk *os.File
	if fileName == "" {
		disk = os.Stdin
	} else {
		disk, err = os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer disk.Close()
	}

	data, err := io.ReadAll(disk)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := tokenize(data)

	if isPrint {
		printContent(content)
	} else if isDir {
		fmt.Println("Listing files...")
	}
}

func tokenize(data []byte) [][]byte {
	var err error
	content := strings.Split(string(data[:len(data)-1]), "\n")

	// Strip header and line prefixes
	content = content[2:]
	for i, line := range content {
		content[i] = line[3:]
	}

	tokens := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		tokens[i] = make([]byte, numBytes/numRows)
	}

	for i, s := range content {
		tokens[i], err = hex.DecodeString(s)
		if err != nil {
			panic(err)
		}
	}

	return tokens
}

func printContent(content [][]byte) {
	fmt.Println(diskHeader)
	for i, row := range content {
		fmt.Printf("%02X:", i)
		for _, chunk := range row {
			fmt.Printf("%02X", chunk)
		}
		fmt.Println()
	}
}
