package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(searchQuery string) {

	res, _ := http.Get("https://amazon.in/s?k="+searchQuery)

	// fmt.Println(res.Status)

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
	log.Fatal(err)
	}

	list := doc.Find("div.s-result-list.s-search-results.sg-row")

	count := 0;

	list.Find("div.a-section.a-spacing-base").Each(func(index int, selector *goquery.Selection) {

	count = count + 1

	var name string
	name = selector.Find("span.a-size-base-plus.a-color-base.a-text-normal").Text()
	var stars string
	stars = selector.Find("span.a-icon-alt").Text()
	var price string
	price = selector.Find("span.a-price-whole").Text()


	if (count <= 5) {
		fmt.Println();
		fmt.Println("Product No. "+strconv.Itoa(count))
		fmt.Println("Product Name: ", name)
		fmt.Println("Rating: ", stars)
		fmt.Println("Price: ₹", price)
	}
})

}

func main() {
	var searchquery string;
	fmt.Println("Enter the keyword to search for")
	fmt.Scanln(&searchquery)
	ExampleScrape(searchquery);
}