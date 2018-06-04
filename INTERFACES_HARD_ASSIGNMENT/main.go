package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// baca argument dari terminal (nama file)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// baca file
	io.Copy(os.Stdout, f)
}
