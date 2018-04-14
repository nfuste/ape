package main

import (
	"fmt"

	"github.com/nfuste/ape/grab"
)

// func sayHello(w http.ResponseWriter, r *http.Request) {

//  fmt.Fprintf(w, `<h1>Hola!</h1>
// <p>Aquí estan les teves cerques:</p>`)

// }

func main() {

	// Primero, el server:
	/*http.HandleFunc("/", sayHello)
	  if err := http.ListenAndServe(":8888", nil); err != nil {
	      fmt.Println(err)
	  }*/

	urls, err := grab.Scraper()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Éxito !!!\n", urls)

	// https://stackoverflow.com/questions/39691100/golang-google-sheets-api-v4-write-update-example
}
