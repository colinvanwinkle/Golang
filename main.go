package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "sync"
  "strings"
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
  "ADSK",
  "AEE",
  "AEP",
  "AES",
  "AET",
  "AFL",
  "AGN",
  "AIG",
  "AIV",
  "AIZ",
  "AJG",
  "AKAM",
  "ALB",
  "ALGN",
  "ALK",
  "ALL",
  "ALLE",
  "ALXN",
  "AMAT",
  "AMD",
  "AME"}

  for _, ticker := range tickers{
    wg.Add(1);
    go func(ticker string){
      doc, _ := goquery.NewDocument("http://www.nasdaq.com/symbol/" + ticker + "/real-time");
      priceChange, _ :=  strconv.ParseFloat(doc.Find("div#qwidget_percent").Contents().Text(), 64);

      if doc.Find(".marginLR10px.arrow-red").Length() == 1{
        wg.Done();
        return;
      }

        tickersFormatTemp := []string{ticker, ":"};
        ticker = strings.Join(tickersFormatTemp, "");
        fmt.Println(ticker, priceChange);

        wg.Done();
      }(ticker)

    }
    wg.Wait();
}
