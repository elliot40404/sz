package main

import (
	"fmt"
	"os"
)

const (
	Help = `
sz version v0.2.0
Prints file sizes in bytes, kilobytes, megabytes and gigabytes
Author: Elliot40404
Usage: sz <file> <file> <file>
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(Help)
		os.Exit(1)
	}
	files := os.Args[1:]
	for _, file := range files {
		printFileSize(file)
	}
}

func printFileSize(path string) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", path)
		return
	}
	if err != nil {
		fmt.Print("Something went wrong\n")
		return
	}
	size := fi.Size()
	fmt.Printf("%d B  %.2f KB  %.3f MB  %.3f GB %s\n",
		size,
		float64(size)/1024,
		float64(size)/1024/1024,
		float64(size)/1024/1024/1024,
		fi.Name(),
	)
}
