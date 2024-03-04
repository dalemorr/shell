package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	AppName      = "shell"
	Version      = "v0.3.0"
	helpUsage    = "display help message"
	versionUsage = "display current version"
	printUsage   = "print raw contents of disk"
	dirUsage     = "print files on disk"
	nameUsage    = "print name of disk"
	catUsage     = "print contents of specified file"
	fileUsage    = "specify name of disk"
)

func main() {
	var err error
	errorLogger := log.New(os.Stderr, "", 0)

	// Parse command
	const (
		isHelp = iota
		isVersion
		isPrint
		isDir
		isName
		isCat
	)

	boolArgs := make([]bool, 6)

	var catFileName string
	var fileName string

	flag.BoolVar(&boolArgs[isHelp], "h", false, helpUsage)
	flag.BoolVar(&boolArgs[isHelp], "?", false, helpUsage)
	flag.BoolVar(&boolArgs[isVersion], "v", false, versionUsage)
	flag.BoolVar(&boolArgs[isPrint], "print", false, printUsage)
	flag.BoolVar(&boolArgs[isDir], "dir", false, dirUsage)
	flag.BoolVar(&boolArgs[isName], "name", false, nameUsage)
	flag.StringVar(&catFileName, "cat", "", catUsage)
	flag.StringVar(&fileName, "f", "", fileUsage)
	flag.Parse()

	if catFileName != "" {
		boolArgs[isCat] = true
	}

	countUniques := 0
	var uniqueIndex int
	for i, arg := range boolArgs {
		if arg {
			uniqueIndex = i
			countUniques++
		}
	}

	if (countUniques != 1) || ((boolArgs[isHelp] || boolArgs[isVersion]) && fileName != "") {
		flag.CommandLine.Usage()
		return
	}

	// Version and help messages
	if boolArgs[isHelp] {
		flag.CommandLine.Usage()
		return
	}
	if boolArgs[isVersion] {
		fmt.Printf("%s %s\n", AppName, Version)
		return
	}

	// Read data
	d := new(disk)
	if fileName == "" {
		err = d.init(os.Stdin)
		if err != nil {
			errorLogger.Fatalln(err)
		}
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			errorLogger.Fatalln(err)
		}
		err = d.init(file)
		if err != nil {
			errorLogger.Fatalln(err)
		}

		defer d.file.Close()
	}

	// Print requested data
	switch uniqueIndex {
	case isPrint:
		d.printFormatted()
	case isDir:
		d.printFiles()
	case isName:
		d.printVolumeName()
	case isCat:
		d.printContent(catFileName, errorLogger)
	default:
		panic(errors.New("an unknown error has occurred"))
	}
}
