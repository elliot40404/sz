package main

import (
	"fmt"
	"os"
)

const (
	VERSION = "0.1"
	AUTHOR  = "Elliot40404"
	DESC	= "Prints the size of a file in bytes, kilobytes, megabytes and gigabytes"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("sz version %s\n", VERSION)
		fmt.Printf("%s\n", DESC)
		fmt.Printf("Author: %s\n", AUTHOR)
		fmt.Printf("Usage: sz <file>\n")
		os.Exit(1)
	}
	// check if the file exists
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", os.Args[1])
		os.Exit(1)
	}
	// print size of the file
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	size := fi.Size()
	// print the size in bytes, kilobytes and megabytes, gigabytes
	fmt.Printf("%d B  ", size)
	fmt.Printf("%.2f KB  ", float64(size)/1024)
	fmt.Printf("%.3f MB  ", float64(size)/1024/1024)
	fmt.Printf("%.3f GB\n", float64(size)/1024/1024/1024)
}
