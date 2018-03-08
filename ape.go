package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func printInfo(i int, s *goquery.Selection) {
	href, _ := s.Attr("href")
	fmt.Printf("%s\n(%s)\n\n", s.Text(), href)
}

func main() {

	doc, err := goquery.NewDocument("http://www.fotocasa.es/es/comprar/casas/castelldefels/zona-platja/l?maxPrice=350000&minRooms=3&minSurface=80")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find(".re-Card-link").Each(printInfo)

}

// mirar https://developers.google.com/sheets/api/quickstart/go?authuser=1
