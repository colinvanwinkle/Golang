package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "sync"
  "strings"
  );

 func getStocksOverOne(ticker string, wg *sync.WaitGroup){
  wg.Add(1);
  doc, _ := goquery.NewDocument("http://www.nasdaq.com/symbol/" + ticker + "/real-time");
  percentChange :=  doc.Find("div#qwidget_percent").Contents().Text();
  percentOneDigit, _ := strconv.Atoi(strings.Split(percentChange, ".")[0]);

  if doc.Find(".marginLR10px.arrow-red").Length() == 1 || percentOneDigit < 1{
    wg.Done();
    return;
  }

    tickersFormatTemp := []string{ticker, ":"};
    ticker = strings.Join(tickersFormatTemp, "");
    fmt.Println(ticker, percentChange);

    wg.Done();
  }


func main(){
  tickers := []string {  "A",  "AAL",  "AAP",  "AAPL",  "ABBV",  "ABC",  "ABT",  "ACN",  "ADBE",  "ADI",  "ADM",  "ADP",  "ADSK",  "AEE",  "AEP",  "AES",  "AET",  "AFL",
     "AGN",  "AIG",  "AIV",  "AIZ",  "AJG",  "AKAM",  "ALB",  "ALGN",  "ALK",  "ALL",  "ALLE",  "ALXN",  "AMAT",  "AMD",  "AME", "AMG", "AMGN", "AMP", "AMT", "AMZN",
      "AN", "ANSS", "ANTM", "AON", "APA", "APC", "APD", "APH", "ARE", "ARNC", "ATVI", "AVB", "AVB", "AVGO", "AVY", "AWK", "AXP", "AYI", "AZO", "BA",
      "BAC", "BAX", "BBBY", "BBT", "BBY", "BCR", "BDX", "BEN"}

  var wg sync.WaitGroup;

  for _, ticker := range tickers{
   go getStocksOverOne(ticker, &wg);
  }
    wg.Wait();
}
