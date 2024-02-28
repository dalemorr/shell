package main

import (
	"encoding/hex"
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
	numRows          = 32
	numColumns       = 64
)

type disk struct {
	file *os.File
	raw  []string
	root *root
}

type root struct {
	index      uint8
	available  *empty
	bad        *bad
	header     *fileHeader
	volumeName string
}

type empty struct {
	index uint8
	next  *empty
}

type bad struct {
	index uint8
	next  *bad
}

type fileHeader struct {
	index uint8
	next  *fileHeader
	data  *fileData
	name  string
}

type fileData struct {
	index uint8
	next  *fileData
	data  string
}

func (d *disk) init(file *os.File) error {
	var err error

	d.file = file
	data, err := io.ReadAll(d.file)
	if err != nil {
		return err
	}

	if !isValid(data) {
		return errors.New(diskErrorMessage)
	}

	d.initRaw(data)
	d.initAvailable()
	d.initBad()
	d.initHeader()
	d.initVolumeName()

	return nil
}

func (d *disk) initRaw(data []byte) {
	content := strings.Split(string(data[:len(data)-1]), "\n")
	content = content[2:]
	for i, line := range content {
		content[i] = line[3:]
	}

	d.raw = content
	d.root = new(root)
	d.root.index = 0
}

func (d *disk) initAvailable() {
	temp, err := hex.DecodeString(d.raw[0][1:3])
	if err != nil {
		panic(err)
	}
	index := temp[0]

	if index != 0 {
		d.root.available = &empty{index, nil}
		currentAvailable := d.root.available
		temp, err := hex.DecodeString(d.raw[index][1:3])
		if err != nil {
			panic(err)
		}
		index := temp[0]

		for index != 0 {
			currentAvailable.next = &empty{index, nil}
			currentAvailable = currentAvailable.next

			temp, err = hex.DecodeString(d.raw[index][1:3])
			if err != nil {
				panic(err)
			}
			index = temp[0]
		}
	} else {
		d.root.available = nil
	}
}

func (d *disk) initBad() {
	var err error
	var temp []byte
	var index uint8

	temp, err = hex.DecodeString(d.raw[0][3:5])
	if err != nil {
		panic(err)
	}
	index = temp[0]

	if index != 0 {
		d.root.bad = &bad{index, nil}
		currentBad := d.root.bad
		temp, err := hex.DecodeString(d.raw[index][1:3])
		if err != nil {
			panic(err)
		}
		index := temp[0]

		for index != 0 {
			currentBad.next = &bad{index, nil}
			currentBad = currentBad.next

			temp, err = hex.DecodeString(d.raw[index][1:3])
			if err != nil {
				panic(err)
			}
			index = temp[0]
		}
	} else {
		d.root.available = nil
	}
}

func (d *disk) initHeader() {
	var err error
	var temp []byte
	var index, index1 uint8

	temp, err = hex.DecodeString(d.raw[0][5:7])
	if err != nil {
		panic(err)
	}
	index = temp[0]

	if index != 0 {
		d.root.header = new(fileHeader)
		d.root.header.index = 0
		currentHeader := d.root.header

		temp, err = hex.DecodeString(d.raw[index][3:5])
		if err != nil {
			panic(err)
		}
		index1 = temp[0]

		if index1 != 0 {
			currentHeader.data = new(fileData)
			currentData := currentHeader.data

			temp, err := hex.DecodeString(d.raw[index1][1:3])
			if err != nil {
				panic(err)
			}
			index1 := temp[0]

			for index1 != 0 {
				currentData.next = new(fileData)
				currentData.index = index1
				temp, err = hex.DecodeString(d.raw[index][3 : numColumns-1])
				if err != nil {
					panic(err)
				}
				currentData.data = string(temp)

				currentData = currentData.next
				temp, err = hex.DecodeString(d.raw[index1][1:3])
				if err != nil {
					panic(err)
				}
				index = temp[0]
			}
		}

		temp, err := hex.DecodeString(d.raw[index][1:3])
		if err != nil {
			panic(err)
		}
		index := temp[0]

		for index != 0 {
			currentHeader.next = new(fileHeader)
			currentHeader.index = index
			temp, err = hex.DecodeString(d.raw[index][3 : numColumns-1])
			if err != nil {
				panic(err)
			}
			currentHeader.name = string(temp)

			currentHeader = currentHeader.next
			temp, err = hex.DecodeString(d.raw[index][1:3])
			if err != nil {
				panic(err)
			}
			index = temp[0]
		}
	} else {
		d.root.available = nil
	}
}

func (d *disk) initVolumeName() {
	var err error
	var temp []byte

	temp, err = hex.DecodeString(d.raw[0][7 : numColumns-1])
	if err != nil {
		panic(err)
	}
	d.root.volumeName = string(temp)
}

func (d *disk) printFormatted() {
	fmt.Println(diskHeader)
	for i, s := range d.raw {
		fmt.Printf("%02X:%s\n", i, s)
	}
}

func (d *disk) printVolumeName() {
	fmt.Println(d.root.volumeName)
}

func (d *disk) printFiles() {
	currentHeader := d.root.header
	for currentHeader != nil {
		fmt.Println(currentHeader.name)
		currentHeader = currentHeader.next
	}
}

func (d *disk) String() string {
	return strings.Join(d.raw, "\n")
}

func isValid(data []byte) bool {
	return true
}
