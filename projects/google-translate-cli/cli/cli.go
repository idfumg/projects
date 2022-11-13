package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jeffail/gabs"
)

type Translator struct {
	SourceLanguage string
	TargetLanguage string
	SourceText     string
}

const url = "https://translate.googleapis.com/translate_a/single"

func NewTranslator(source, target, text string) *Translator {
	return &Translator{
		SourceLanguage: source,
		TargetLanguage: target,
		SourceText:     text,
	}
}

func (t *Translator) Translate() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", t.SourceLanguage)
	query.Add("tl", t.TargetLanguage)
	query.Add("dt", "t")
	query.Add("q", t.SourceText)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("You have been rate limited. Try again later")
		os.Exit(1)
	}

	parsed, err := gabs.ParseJSONBuffer(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}

	nestOne, err := parsed.ArrayElement(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}
	
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}

	translated, err := nestTwo.ArrayElement(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}
	
	data := translated.Data().(string)
	fmt.Println(data)
}
