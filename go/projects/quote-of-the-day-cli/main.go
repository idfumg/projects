package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	response, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}