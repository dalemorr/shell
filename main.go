package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var err error
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

	fmt.Printf("%#v\n", contentArray)
}
