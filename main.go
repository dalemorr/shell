package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var err error
	var disk *os.File
	version := "0.0.1"

	versionFlag := flag.Bool("v", false, "display current version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("shell v%s\n", version)
		return
	}

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

	fmt.Println("XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF")
	for i, line := range contentArray {
		if i < 0x10 {
			fmt.Printf("0%X:%s\n", i, line)
		} else {
			fmt.Printf("%X:%s\n", i, line)
		}
	}
}
