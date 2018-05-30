package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// resp, err := http.Get("http://id.techinasia.com/")
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// byte slice from resp
	// 99999 = number of element avilable
	bs := make([]byte, 99999)

}
