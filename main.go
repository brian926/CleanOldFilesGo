package main

import (
	"io/ioutil"
	"os"
	"time"
	"fmt"
)

func isOlderThan(t time.Time) bool {
	d := 24 * time.Hour
	return time.Now().Sub(t) > d * 30
}

func findFiles(dir string) (files []os.FileInfo, err error) {
	tmpfiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range tmpfiles {
		if file.Mode().IsRegular() {
			if isOlderThan(file.ModTime()) {
				files = append(files, file)
				//fmt.Println(dir + "/" + file.Name())
				e := os.Remove(dir + "/" + file.Name())
				if e != nil {
					fmt.Println(e)
				  } 
			}
		}
	}
	return
}

func main() {
	data := os.Args[1:]

	for _, args := range data {
		findFiles(args)
	}
}