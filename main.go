package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func listDir(dir string, limit int) {
	// Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// Print all the file names
	for idx, file := range files {
		if idx > limit {
			return
		}
		fmt.Println(file.Name())
	}
}

func cat(file string) {
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	fmt.Println("==> Root Dir")
	listDir("/", 10)
	fmt.Println("==> Proc Dir")
	listDir("/proc", 10)
	fmt.Println("==> PID 1 Commandline Arguments")
	cat("/proc/1/cmdline")

}
