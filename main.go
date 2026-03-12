
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)
type Stock struct{
	name string
	price string
	change string
}
func main() {
	items := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"COST",
		"MCD",
		"V",
		"WMT",
		"DIS",
		"MMM",
		"INTC",
  		"AXP",
    	"AAPL",
    	"BA",
    	"CSCO",
    	"GS",
    	"JPM",
    	"CRM",
    	"VZ",
	}
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 3})
	stocks := []Stock{}
	c.OnHTML("div#nimbus-app", func(e *colly.HTMLElement) {
		
		stock := Stock{}

		stock.name = e.ChildText("h1")
		// fmt.Println("Company:", stock.name)
		stock.price = e.ChildText("span[data-testid='qsp-price']")
		// fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("span[data-testid='qsp-price-change-percent']")
		// fmt.Println("Change:", stock.change)
		if stock.name == "" || stock.price == "" || stock.change == ""{
			fmt.Println(e.Request.URL)
			return
		}
		stocks = append(stocks, stock)
	})
	domain := "https://finance.yahoo.com/quote/"
	for _,val := range items{
		temp := domain+val+"/"
		c.Visit(temp)
	}
	c.Wait()
	fmt.Println(len(stocks))
	
}