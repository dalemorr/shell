package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	AppName = "shell"
	Version = "v0.1.0"
)

func main() {
	var err error

	// Handle args
	hFlag1 := flag.Bool("h", false, "display help message")
	hFlag2 := flag.Bool("H", false, "display help message")
	hFlag3 := flag.Bool("?", false, "display help message")
	vFlag1 := flag.Bool("v", false, "display current version")
	vFlag2 := flag.Bool("V", false, "display current version")
	printFlag := flag.Bool("print", false, "print raw contents of disk")
	dirFlag := flag.Bool("dir", false, "list files on the disk")
	// fileFlag := flag.String("f", "", "specify file name of disk")
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
	if len(os.Args) > 1 {
		disk, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		defer disk.Close()
	} else {
		disk = os.Stdin
	}

	data, err := io.ReadAll(disk)
	if err != nil {
		fmt.Println(err)
	}

	content := string(data[:len(data)-1])
	contentArray := strings.Split(content, "\n")
	contentArray = contentArray[2:]
	for i, line := range contentArray {
		contentArray[i] = line[3:]
	}

	if *printFlag {
		printFormattedContent(contentArray)
	} else if *dirFlag {
		fmt.Println("..welp")
	}
}

func printFormattedContent(content []string) {
	fmt.Println("XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF")
	for i, line := range content {
		if i < 0x10 {
			fmt.Printf("0%X:%s\n", i, line)
		} else {
			fmt.Printf("%X:%s\n", i, line)
		}
	}
}
