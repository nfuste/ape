package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hola!</h1>
<p>Aquí estan les teves cerques:</p>`)

}

func main() {

	//Primero, el server:
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}

}
