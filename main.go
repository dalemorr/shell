package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	AppName      = "shell"
	Version      = "v0.2.1"
	helpUsage    = "display help message"
	versionUsage = "display current version"
	printUsage   = "print raw contents of disk"
	dirUsage     = "list files on disk"
	fileUsage    = "specify name of disk"
)

func main() {
	var err error

	// Parse command
	var isHelp bool
	var isVersion bool
	var isPrint bool
	var isDir bool
	var fileName string

	flag.BoolVar(&isHelp, "h", false, helpUsage)
	flag.BoolVar(&isHelp, "?", false, helpUsage)
	flag.BoolVar(&isVersion, "v", false, versionUsage)
	flag.BoolVar(&isPrint, "print", false, printUsage)
	flag.BoolVar(&isDir, "dir", false, dirUsage)
	flag.StringVar(&fileName, "f", "", fileUsage)
	flag.Parse()

	countUniques := 0
	if isHelp {
		countUniques++
	}
	if isVersion {
		countUniques++
	}
	if isPrint {
		countUniques++
	}
	if isDir {
		countUniques++
	}

	if (countUniques != 1) || ((isHelp || isVersion) && fileName != "") {
		flag.CommandLine.Usage()
		return
	}

	// Version and help messages
	if isHelp {
		flag.CommandLine.Usage()
		return
	}
	if isVersion {
		fmt.Printf("%s %s\n", AppName, Version)
		return
	}

	// Read data
	d := new(disk)
	if fileName == "" {
		d.file = os.Stdin
	} else {
		d.file, err = os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer d.file.Close()
	}

	// Print requested data
	if isPrint {
		d.printRaw()
	} else if isDir {
		d.printFiles()
	}
}

func isValid(data []byte) bool {
	return true
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
		tokens[i] = make([]byte, numColumns)
	}

	for i, line := range content {
		tokens[i], err = decodeOctal(line)
		if err != nil {
			panic(err)
		}
	}

	return tokens
}
