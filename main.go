package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
);

func main(){
  tickers := []string {
  "A",
  "AAL",
  "AAP",
  "AAPL",
  "ABBV",
  "ABC",
  "ABT",
  "ACN",
  "ADBE",
  "ADI",
  "ADM",
  "ADP",
  "ADSK" }

  for _, ticker := range tickers{
    doc, _ := goquery.NewDocument("http://www.nasdaq.com/symbol/" + ticker + "/real-time");
    priceChange, _ :=  strconv.ParseFloat(doc.Find("div#qwidget_netchange").Contents().Text(), 64);

    if doc.Find(".marginLR10px.arrow-red").Length() == 1{
      priceChange = -1 * priceChange;
    }
      fmt.Println(priceChange);

    }
}
