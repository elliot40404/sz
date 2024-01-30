package main

import (
	"fmt"
	"os"
)

const (
	Help = `
	sz version v0.1.0
	Prints the size of a file in bytes, kilobytes, megabytes and gigabytes
	Author: Elliot40404
	Usage: sz <file>
	`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(Help)
		os.Exit(1)
	}
	// check if the file exists
	fi, err := os.Stat(os.Args[1])
	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", os.Args[1])
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	size := fi.Size()
	// print the size in bytes, kilobytes and megabytes, gigabytes
	fmt.Printf("%d B  %.2f KB  %.3f MB  %.3f GB\n",
		size,
		float64(size)/1024,
		float64(size)/1024/1024,
		float64(size)/1024/1024/1024,
	)
}
