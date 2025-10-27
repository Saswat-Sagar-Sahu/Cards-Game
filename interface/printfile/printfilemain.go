package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := os.Args[1]
	f, err := os.Open(s)
	if(err != nil){
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
