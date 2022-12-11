package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	round := rps.PlayRound(playerChoice)
	log.Println(round.Message, round.ComputerChoice, round.RoundResult)
	out, err := json.MarshalIndent(round, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	fmt.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}