package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ficha struct {
	zona   float64
	precio float64
}

func main() {
	url := "http://fotocasa.es"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		fmt.Println(err)
		return
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)

	// Test strings
	fmt.Println(strings.Contains("Vivo en Toledo", "Toledo"))
}
