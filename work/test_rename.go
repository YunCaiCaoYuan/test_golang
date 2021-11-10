package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main()  {
	path := "/Users/sunbin/work/my_doc/2"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for index, f := range files {
		fmt.Println()
		err := os.Rename(path + "/" + f.Name(), path + "/" + fmt.Sprintf("head_%d.jpg", index))
		if err != nil {
			fmt.Println("filename:", f.Name(), "err:", err)
		}
	}
}
