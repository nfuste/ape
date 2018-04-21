package main

import (
	"fmt"
	"net/http"

	"github.com/nfuste/ape/grab"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	urls, err := grab.Scraper()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Éxito !!!\n", urls)

}

func main() {

	// mux es una tabla de correspondencias entre los paths y los handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
