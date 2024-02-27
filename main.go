package main

import (
	"errors"
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
	dirUsage     = "list files on disk"
	fileUsage    = "specify name of disk"
	diskHeader   = "XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	numRows    = 32
	numColumns = 64
	numNibbles = numRows * numColumns
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

	fmt.Printf("%d\n", content[9][2])

	// Print requested data
	if isPrint {
		printDisk(content)
	} else if isDir {
		printFiles(content)
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

func printDisk(content [][]byte) {
	fmt.Println(diskHeader)
	for i, row := range content {
		fmt.Printf("%02X:", i)
		for _, nibble := range row {
			fmt.Printf("%X", nibble)
		}
		fmt.Println()
	}
}

func printFiles(content [][]byte) {
	fmt.Println("Listing files...")
}

func decodeOctal(s string) ([]byte, error) {
	if !isOctal(s) {
		return []byte{}, errors.New("invalid octal string")
	}

	data := make([]byte, len(s))

	for i, char := range s {
		if char >= 0x30 && char <= 0x39 {
			data[i] = byte(char - 0x30)
		} else if char >= 0x61 && char <= 0x66 {
			data[i] = byte(char - 0x51)
		}
	}

	return data, nil
}

func isOctal(s string) bool {
	str := strings.ToLower(s)

	for _, char := range str {
		if !((char >= 0x30 && char <= 0x39) || (char >= 0x61 && char <= 0x66)) {
			return false
		}
	}

	return true
}
