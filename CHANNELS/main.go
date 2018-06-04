package main

import (
	"fmt"
	"net/http"
	"time"
)

// mau cek apakah websitenya = up / down
func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// bikin channels
	c := make(chan string)

	for _, site := range websites {
		go checkWebsite(site, c)
	}

	for s := range c {
		// // s = <-c
		// go checkWebsite(s, c)
		// functin literal / anonymouse func / lambda
		go func(site string) {
			// kasih delay 5detik, biar ga spamming
			time.Sleep(5 * time.Second)
			checkWebsite(site, c)
		}(s) // tutup kurung () buat self call func ini
	}

}

// argument kedua = c dengan type channel, return value string
func checkWebsite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site + " mungkin lagi mati.")
		c <- site
		return
	}

	fmt.Println(site + " baik-baik aja.")
	c <- site
}
