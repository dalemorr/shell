package main

import (
	"errors"
	"io"
	"os"
)

const (
	diskHeader = "XX:                1               2               3\n" +
		"XX:0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	numRows    = 32
	numColumns = 64
	numNibbles = numRows * numColumns
)

type disk struct {
	file *os.File
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
		return errors.New("invalid disk")
	}

	content := tokenize(data)

	return nil
}

func (d *disk) printRaw() {

}

func (d *disk) printFiles() {

}
