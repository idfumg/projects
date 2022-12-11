package main

import (
	"flag"
	"fmt"
	"os"

	"google-translator/cli"
)

func parseFlags() (string, string, string) {
	var (
		sourceLanguage string
		targetLanguage string
		sourceText     string
	)

	flag.StringVar(&sourceLanguage, "s", "en", "Source language[en]")
	flag.StringVar(&targetLanguage, "t", "ru", "Target language[ru]")
	flag.StringVar(&sourceText, "text", "", "Text to translate")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return sourceLanguage, targetLanguage, sourceText
}

func main() {
	sourceLanguage, targetLanguage, sourceText := parseFlags()
	translator := cli.NewTranslator(sourceLanguage, targetLanguage, sourceText)
	translator.Translate()
}
