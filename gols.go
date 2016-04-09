package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	//Set clolor
	//ex. fmt.Println(red,"text",reset)
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

func LsPrint(path ...string) {
	for _, r := range path {
		//Processing for each argument
		fileinfos, _ := ioutil.ReadDir(r)
		for _, fileinfo := range fileinfos {
			fmt.Println(fileinfo.Name())
		}
	}
}

func main() {

	var (
		//Flags
		a = flag.Bool("a", false, "Show all files")
		f = flag.Bool("f", false, "Not sort")
		m = flag.Bool("m", false, "Format commas")
		l = flag.Bool("l", false, "Show detail")
	)
	flag.Parse()
	//Parameteir args
	args := flag.Args()

	//Get current directory path
	curDir, _ := os.Getwd()
	curDir += "/"

	//If no args, use curDir
	if len(args) == 0 {
		//No argument
		//append(args, curDir)
		//args[0] = curDir
		args := []string{curDir}
		//fmt.Println(args)
		LsPrint(args...)
		//fmt.Println(curDir)
	} else {
		//There arguments
		LsPrint(args...)
		//fmt.Println(green, "args:", args, "len:", len(args), reset)
	}

	fmt.Println("-a:", *a, "-f:", *f, "-m:", *m, "-l:", *l)
}
