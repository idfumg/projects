package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func main() {
	var r RSS

	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Below are all the Google Search Trends for today!")
	fmt.Println("-------------------------------------------------")

	for i := range r.Channel.ItemList {
		fmt.Println("#", i + 1)
		fmt.Println("Search term:", r.Channel.ItemList[i].Title)
		fmt.Println("Link to the trend:", r.Channel.ItemList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("HeadlineLink:", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)
	}
}
