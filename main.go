package main

import (
	"fmt"
	"sync"

	"github.com/bits-and-atoms/stock_scrapper/filemanager"
	"github.com/gocolly/colly"
)

type Stock struct {
	Name   string `json:"company name"`
	Price  string `json:"price"`
	Change string `json:"percentage change"`
}

func main() {
	fm := filemanager.FileManager{}
	fm.Set("items.txt", "stocks.json")
	items := []string{}
	err := fm.ReadLines(&items)
	if err != nil {
		fmt.Println("cant read file")
		return
	}
	
	var stocks []Stock
	var mu sync.Mutex
	
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 20})
	c.OnHTML("div#nimbus-app", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.Name = e.ChildText("h1")
		// fmt.Println("Company:", stock.Name)
		stock.Price = e.ChildText("span[data-testid='qsp-price']")
		// fmt.Println("Price:", stock.Price)
		stock.Change = e.ChildText("span[data-testid='qsp-price-change-percent']")
		// fmt.Println("Change:", stock.Change)
		if stock.Name == "" || stock.Price == "" || stock.Change == "" {
			fmt.Println("site dont exist", e.Request.URL)
			return
		}
		
		mu.Lock()
		stocks = append(stocks, stock)
		mu.Unlock()
	})
	domain := "https://finance.yahoo.com/quote/"

	for _, val := range items {
		temp := domain + val + "/"
		c.Visit(temp)
	}
	c.Wait()
	
	err = fm.WriteJson(stocks)
	if err != nil {
		fmt.Println("cant write to file")
		return
	}
}