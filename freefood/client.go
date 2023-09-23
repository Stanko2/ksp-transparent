package freefood

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func Download() {
	c := colly.NewCollector(colly.AllowedDomains("www.freefood.sk"))
	c.OnHTML("#free-food .day-offer", func(e *colly.HTMLElement) {
		day := e.DOM.Parent().Children().Nodes[0].FirstChild.Data
		fmt.Println(day)
		dayOffers := e.DOM.Children().Nodes
		ret := make([]Food, len(dayOffers))
		for i, dayOffer := range dayOffers {
			ret[i].Name = strings.Split(dayOffer.FirstChild.NextSibling.Data, "A:")[0]
			priceStr := dayOffer.LastChild.FirstChild.Data
			price, err := strconv.ParseFloat(priceStr[0:len(priceStr) - 3], 64)
			if err != nil {
				panic(err)
			}
			ret[i].Price = price
			fmt.Print(dayOffer.FirstChild.FirstChild.Data)
			fmt.Print(" - ")
			fmt.Print(ret[i].Name)
			fmt.Print(" - ")
			fmt.Println(price)
		}
		// day := e.DOM.Parent().Children().Find(".day-title").Text()
		// fmt.Println(day)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("http://www.freefood.sk/menu/")
	if err != nil {
		panic(err)
	}
}