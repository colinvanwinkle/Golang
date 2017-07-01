package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "sync"
  "strings"
  "time"
  );

 /*Given a stock ticker, ticker, this function
  *determines if it was over a 1% gain, and if
  *so, prints it to the output
  */
 func getStocksOverOne(ticker string, wg *sync.WaitGroup){
  wg.Add(1);
  //gets html
  doc, _ := goquery.NewDocument("http://www.nasdaq.com/symbol/" + ticker + "/real-time");
  //finds element that has the % change
  percentChange :=  doc.Find("div#qwidget_percent").Contents().Text();
  percentOneDigit, _ := strconv.Atoi(strings.Split(percentChange, ".")[0]);

  //quits this iteration if it less than 1% gain
  if doc.Find(".marginLR10px.arrow-red").Length() == 1 || percentOneDigit < 1{
    wg.Done();
    return;
  }

    //formats and prints to console if it is over 1%
    tickersFormatTemp := []string{ticker, ":"};
    ticker = strings.Join(tickersFormatTemp, "");
    fmt.Println(ticker, percentChange);

    wg.Done();
  }


func main(){
  //tickers to query
  tickers := []string {  "A",  "AAL",  "AAP",  "AAPL",  "ABBV",  "ABC",  "ABT",  "ACN",  "ADBE",  "ADI",  "ADM",  "ADP",  "ADSK",  "AEE",  "AEP",  "AES",  "AET",  "AFL",
     "AGN",  "AIG",  "AIV",  "AIZ",  "AJG",  "AKAM",  "ALB",  "ALGN",  "ALK",  "ALL",  "ALLE",  "ALXN",  "AMAT",  "AMD",  "AME", "AMG", "AMGN", "AMP", "AMT", "AMZN",
      "AN", "ANSS", "ANTM", "AON", "APA", "APC", "APD", "APH", "ARE", "ARNC", "ATVI", "AVB", "AVB", "AVGO", "AVY", "AWK", "AXP", "AYI", "AZO", "BA",
      "BAC", "BAX", "BBBY", "BBT", "BBY", "BCR", "BDX", "BEN"}

  var wg sync.WaitGroup;

  for _, ticker := range tickers{
   //looks like i've been getting blocked from sending http requests, put a second delay on each request
   time.Sleep(time.Second)

   go getStocksOverOne(ticker, &wg);
  }
    wg.Wait();
}
