# stock_scrapper

A Go-based web scraper that fetches real-time stock information from Yahoo Finance.

## Features

- Scrapes stock data including company name, current price, and percentage change
- Supports concurrent scraping of multiple stocks
- Outputs data in JSON format

## Usage

1. Add stock ticker symbols to `items.txt`, one per line:
   ```
   AAPL
   MSFT
   GOOGL
   ```

2. Run the scraper:
   ```bash
   go run .
   ```

3. Check the output in `stocks.json` - it will contain an array of stock data with company name, price, and percentage change for each ticker.

## Dependencies

- [gocolly/colly](https://github.com/gocolly/colly) - Web scraping framework

Install dependencies with:
```bash
go mod download
```
