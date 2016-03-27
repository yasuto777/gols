package main

import (
	"fmt"
	//"io/ioutil"
	"flag"
	"os"
)

const (
	//Set clolor
	black   = "\x1b[30m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"

	under_bar = "\x1b[4m"
	reset     = "\x1b[0m"
)

//func IsDir(name string) (isDir bool, err os.Error) {
//	finfo, err := os.Stat(name)
//	if err != nil {
//		return false, err
//	}
//	return finfo.IsDirectory(), nil
//}

func main() {

	var (
		//Flags
		a = flag.Bool("a", false, "Show all files")
		f = flag.Bool("f", false, "Not sort")
		m = flag.Bool("m", false, "Format commas")
		l = flag.Bool("l", false, "Show detail")

		//Parameteir args
		args string
	)
	flag.Parse()

	//Get current directory path
	curDir, _ := os.Getwd()
	curDir += "/"

	//If no args, use curDir
	if arg == "" {
		arg = curDir
	}
	//fmt.Println(*a, *l)
}
