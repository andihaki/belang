package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// resp, err := http.Get("http://id.techinasia.com/")
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// // byte slice from resp
	// // 99999 = number of element avilable
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// fmt.Print(resp.ContentLength)

	lw := logWriter{}

	// lw, sebelumnya = os.stdout
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Total bytes: ", len(bs))
	return len(bs), nil
}
