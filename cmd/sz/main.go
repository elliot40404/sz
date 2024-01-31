package main

import (
	"fmt"
	"os"
)

const (
	Help = `
sz version v0.2.0
Prints file sizes in bytes, kilobytes, megabytes, and gigabytes
Author: Elliot40404
Usage: sz <file> <file> <file>
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(Help)
		os.Exit(1)
	}

	for _, file := range os.Args[1:] {
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
	sizeKb := float64(size) / 1024
	sizeMb := sizeKb / 1024
	sizeGb := sizeMb / 1024
	fmt.Printf("%d B  %.2f KB  %.3f MB  %.3f GB %s\n",
		size,
		sizeKb,
		sizeMb,
		sizeGb,
		fi.Name(),
	)
}
