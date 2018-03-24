package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func printInfo(i int, s *goquery.Selection) {
	href, _ := s.Attr("href")
	fmt.Printf("%s\n(%s)\n\n", s.Text(), href)
}

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

	// Apartado Fotocasa: aquí cogeremos todos los inmuebles de la búsqueda de Fotocasa

	fmt.Println("Aquí están los inmuebles de Fotocasa:")

	doc, err := goquery.NewDocument("http://www.fotocasa.es/es/comprar/casas/castelldefels/zona-platja/l?maxPrice=350000&minRooms=3&minSurface=80")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find(".re-Card-link").Each(printInfo)

	// Apartado Idealista: aquí cogeremos todos los inmuebles de la búsqueda de Idealista
	fmt.Println("Aquí están los de Idealista:")

	doc2, err := goquery.NewDocument("https://www.idealista.com/areas/venta-viviendas/con-precio-hasta_360000,chalets,casas-de-pueblo,duplex,aticos,de-tres-dormitorios,de-cuatro-cinco-habitaciones-o-mas,dos-banos,tres-banos-o-mas/?shape=%28%28efzzFwjyJuHwByE%7D_A%7By%40zKiNwn%40r%5B_N~Y%7Bp%40rDuPkZq%5C%7DOyw%40yToy%40zu%40cHlEyJe%40moApYr%40rBloDfDh%7DF%29%29")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc2.Find(".item-info-container").Each(printInfo)

	// https://stackoverflow.com/questions/39691100/golang-google-sheets-api-v4-write-update-example
}
