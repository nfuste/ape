package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey there!</h1>
<p>This is a test</p>`)
	fmt.Printf(Data)

}

func main() {

	//Primero, el server:
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}

}
