package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/oauth2"
)

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

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
