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
	Version      = "v0.2.2"
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
	var isHelp bool
	var isVersion bool
	var isPrint bool
	var isDir bool
	var isName bool
	var catFileName string
	var fileName string

	flag.BoolVar(&isHelp, "h", false, helpUsage)
	flag.BoolVar(&isHelp, "?", false, helpUsage)
	flag.BoolVar(&isVersion, "v", false, versionUsage)
	flag.BoolVar(&isPrint, "print", false, printUsage)
	flag.BoolVar(&isDir, "dir", false, dirUsage)
	flag.BoolVar(&isName, "name", false, nameUsage)
	flag.StringVar(&catFileName, "cat", "", catUsage)
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
	if isName {
		countUniques++
	}
	if catFileName != "" {
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
	if isPrint {
		d.printFormatted()
	} else if isDir {
		d.printFiles()
	} else if isName {
		d.printVolumeName()
	} else if catFileName != "" {
		d.printContent(catFileName, errorLogger)
	} else {
		panic(errors.New("an unknown error has occurred"))
	}
}
