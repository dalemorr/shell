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
	AppName        = "shell"
	Version        = "v0.1.0"
	helpMessage    = "display help message"
	versionMessage = "display current version"
	printMessage   = "print raw contents of disk"
	dirMessage     = "list files on the disk"
	fileFlag       = "specify name of disk"
	diskHeader     = "XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	numRows    = 32
	numColumns = 64
	numBytes   = numRows * numColumns / 2
)

func main() {
	var err error

	// Handle args
	hFlag1 := flag.Bool("h", false, helpMessage)
	hFlag2 := flag.Bool("H", false, helpMessage)
	hFlag3 := flag.Bool("?", false, helpMessage)
	vFlag1 := flag.Bool("v", false, versionMessage)
	vFlag2 := flag.Bool("V", false, versionMessage)
	printFlag := flag.Bool("print", false, printMessage)
	dirFlag := flag.Bool("dir", false, dirMessage)
	fileFlag := flag.String("f", "", fileFlag)
	flag.Parse()

	uniqueFlags := []*bool{hFlag1, hFlag2, hFlag3, vFlag1, vFlag2, printFlag, dirFlag}
	hasUniqueFlag := false
	for _, uFlag := range uniqueFlags {
		if *uFlag {
			if hasUniqueFlag {
				flag.CommandLine.Usage()
				return
			} else {
				hasUniqueFlag = true
			}
		}
	}

	hFlag := new(bool)
	*hFlag = *hFlag1 || *hFlag2 || *hFlag3
	vFlag := new(bool)
	*vFlag = *vFlag1 || *vFlag2

	if *vFlag {
		fmt.Printf("%s %s\n", AppName, Version)
		return
	}
	if *hFlag || !hasUniqueFlag {
		flag.CommandLine.Usage()
		return
	}

	// Read data
	var disk *os.File
	if *fileFlag == "" {
		disk = os.Stdin
	} else {
		disk, err = os.Open(*fileFlag)
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

	content := Tokenize(data)

	if *printFlag {
		PrintContent(content)
	} else if *dirFlag {
		fmt.Println("Listing files...")
	}
}

func Tokenize(data []byte) [][]byte {
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

func PrintContent(content [][]byte) {
	fmt.Println(diskHeader)
	for i, row := range content {
		fmt.Printf("%02X:", i)
		for _, chunk := range row {
			fmt.Printf("%02X", chunk)
		}
		fmt.Println()
	}
}
