package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

func main() {
	resp, err := http.Get("http://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=en")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var res Quote
	if err := json.Unmarshal(body, &res); err != nil {
		os.Exit(1)
	}

	fmt.Println(res.QuoteText + " - " + res.QuoteAuthor)
}
