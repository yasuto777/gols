package main

import "fmt"
import "os"
import "io/ioutil"

const (
	black   = "\x1b[30m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"
	reset   = "\x1b[0m"

	under_bar = "\x1b[4m"
)

func IsDirectory(name string) (isDir bool, err error) {
	//Return fileinfo
	finfo, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	//Check Directory or File
	return finfo.IsDir(), nil
}

func main() {
	// Get current path
	current_dir, _ := os.Getwd()
	// Get files & directories name
	files, _ := ioutil.ReadDir(current_dir)

	var (
		// Most longest file name
		max_len int
		// Tab
		tab string = "\t"
		f   []string
		dir []string
		obj bool
	)

	for _, file := range files {
		//Search long file name.
		if max_len < len(file.Name()) {
			max_len = len(file.Name())
		}
		//IsDirectory
		obj, _ = IsDirectory(file.Name())
		switch obj {
		case true:
			append(dir, file.Name())
		default:
			append(f, file.Name())
		}
	}
	//fmt.Printf("%s%s%s\n",cyan,dir,reset)
	//fmt.Printf("%s\n",f)
	fmt.Println(dir, f)
	fmt.Printf("%s%d%s %stab:%d%s\n", red, max_len, reset, under_bar, len(tab), reset)
}
