package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	AppName      = "shell"
	Version      = "v0.2.1"
	helpUsage    = "display help message"
	versionUsage = "display current version"
	printUsage   = "print raw contents of disk"
	dirUsage     = "list files on disk"
	fileUsage    = "specify name of disk"
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
	d := new(disk)
	if fileName == "" {
		err = d.init(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = d.init(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer d.file.Close()
	}

	// Print requested data
	if isPrint {
		d.printFormatted()
	} else if isDir {
		d.printFiles()
	}
}
