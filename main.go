package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "sync"
);

func main(){

  var wg sync.WaitGroup;

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

    wg.Add(1);
    go func(ticker string){
      doc, _ := goquery.NewDocument("http://www.nasdaq.com/symbol/" + ticker + "/real-time");
      priceChange, _ :=  strconv.ParseFloat(doc.Find("div#qwidget_netchange").Contents().Text(), 64);

      if doc.Find(".marginLR10px.arrow-red").Length() == 1{
        priceChange = -1 * priceChange;
      }
        fmt.Println(priceChange);
        wg.Done();
      }(ticker)

    }
    wg.Wait();
}
