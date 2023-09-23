package freefood

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

func Parse(dayOffers []*html.Node) []Food {
	ret := make([]Food, len(dayOffers))
	for i, dayOffer := range dayOffers {
		ret[i].Name = strings.Split(dayOffer.FirstChild.NextSibling.Data, "A:")[0]
		priceStr := dayOffer.LastChild.FirstChild.Data
		price, err := strconv.ParseFloat(priceStr[0:len(priceStr) - 3], 64)
		if err != nil {
			panic(err)
		}
		ret[i].Price = price
		if (dayOffer.FirstChild.FirstChild.Data == "P. ") {
			ret[i].Type = Soup
		} else {
			ret[i].Type = Main
		}
		
		// fmt.Print(dayOffer.FirstChild.FirstChild.Data)
		// fmt.Print(" - ")
		// fmt.Print(ret[i].Name)
		// fmt.Print(" - ")
		// fmt.Println(price)
	}

	return ret
}

func Atoi(param string) int {
    n, _ := strconv.Atoi(param)
    return n
}

func ParseRestaurant(root *colly.HTMLElement) DailyMenu {
	day := strings.Split(root.DOM.Parent().Children().Nodes[0].FirstChild.Data, ", ")[1]
	dateStr := strings.Split(day, ".")
	date := time.Date(
		Atoi(dateStr[2]),
		time.Month(Atoi(dateStr[1])),
		Atoi(dateStr[0]),
		 0, 0, 0, 0, time.Local)
	menus := Parse(root.DOM.Children().Nodes)
	return DailyMenu{menus: menus, date: date}
}

func Download() map[Restaurant][]DailyMenu {
	ret := make(map[Restaurant][]DailyMenu)
	c := colly.NewCollector(colly.AllowedDomains("www.freefood.sk"))
	c.OnHTML("#fayn-food .day-offer", func(e *colly.HTMLElement) {
		ret[faynfood] = append(ret[faynfood], ParseRestaurant(e))
	})

	c.OnHTML("#free-food .day-offer", func(e *colly.HTMLElement) {
		ret[freefood] = append(ret[freefood], ParseRestaurant(e))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("http://www.freefood.sk/menu/")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(ret)
	return ret
}

func first(i int, err error) {
	panic("unimplemented")
}