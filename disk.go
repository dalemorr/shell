package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	diskHeader = "XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	diskErrorMessage = "invalid disk"
)

type disk struct {
	file *os.File
	raw  []string
	r    *root
}

type root struct {
	available  *empty
	bad        *bad
	header     *fileHeader
	volumeName string
}

type empty struct {
	next *empty
}

type bad struct {
	next *bad
}

type fileHeader struct {
	next *fileHeader
	data *fileData
	name string
}

type fileData struct {
	next *fileData
	data string
}

func (d *disk) init(file *os.File) error {
	d.file = file
	data, err := io.ReadAll(d.file)
	if err != nil {
		return err
	}

	if !isValid(data) {
		return errors.New(diskErrorMessage)
	}

	content := strings.Split(string(data[:len(data)-1]), "\n")
	content = content[2:]
	for i, line := range content {
		content[i] = line[3:]
	}

	d.raw = content

	return nil
}

func (d *disk) printFormatted() {
	fmt.Println(diskHeader)
	for i, s := range d.raw {
		fmt.Printf("%02X:%s\n", i, s)
	}
}

func (d *disk) printFiles() {

}

func (d *disk) String() string {
	return strings.Join(d.raw, "\n")
}

func isValid(data []byte) bool {
	return true
}
