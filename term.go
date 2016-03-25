package main

import(
		"fmt"
		"os"
		"io/ioutil"
		"github.com/nsf/termbox-go"
		)

func main(){
		err := termbox.Init()
		if err != nil{
		panic(err)
		}
		//w:width,h:height
		//w,h:=termbox.Size()
		//termbox.Close()

		current_dir,_:=os.Getwd()
		fileinfos,_:=ioutil.ReadDir(current_dir)

		for _,fileinfo := range fileinfos{
		fmt.Print(fileinfo.Name())
		}
}
