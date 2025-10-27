package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main(){
	resp, err := http.Get("https://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body) //same as previous 3 lines

	lw := logwriter{}
	io.Copy(lw, resp.Body)
}


func (logwriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	return len(bs), nil
}