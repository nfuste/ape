package grab

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// printWrapper coge una lista de strings y le añade la url
// que encuentra en el Selection dado
func printWrapper(urls []string) func(int, *goquery.Selection) {
	return func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Printf("%s\n(%s)\n\n", s.Text(), url)
		urls = append(urls, url)
	}
}

func Scraper() ([]string, error) {
	// aqui vas guardant les urls
	savedUrls := []string{}

	// Apartado Fotocasa: aquí cogeremos todos los inmuebles de la búsqueda de Fotocasa

	fmt.Println("Aquí están los inmuebles de Fotocasa:")

	doc, err := goquery.NewDocument("http://www.fotocasa.es/es/comprar/casas/castelldefels/zona-platja/l?maxPrice=350000&minRooms=3&minSurface=80")
	if err != nil {
		return nil, err
	}

	doc.Find(".re-Card-link").Each(printWrapper(savedUrls))

	// Apartado Idealista: aquí cogeremos todos los inmuebles de la búsqueda de Idealista
	fmt.Println("Aquí están los de Idealista:")

	doc2, err := goquery.NewDocument("https://www.idealista.com/areas/venta-viviendas/con-precio-hasta_360000,chalets,casas-de-pueblo,duplex,aticos,de-tres-dormitorios,de-cuatro-cinco-habitaciones-o-mas,dos-banos,tres-banos-o-mas/?shape=%28%28efzzFwjyJuHwByE%7D_A%7By%40zKiNwn%40r%5B_N~Y%7Bp%40rDuPkZq%5C%7DOyw%40yToy%40zu%40cHlEyJe%40moApYr%40rBloDfDh%7DF%29%29")
	if err != nil {
		return nil, err
	}

	doc2.Find(".item-info-container").Each(printWrapper(savedUrls))

	return savedUrls, nil
}
