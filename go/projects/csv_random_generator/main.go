package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	delimiter = ";"
	colsSize  = 6
	rowsSize  = 20
	wordSize  = 6
	quote     = ""
	alphabet  = "0123456789abcdefghklmnoprstuwxyzABCDEFGHKLMNOPRSTUWXYZ"
)

func GenerateWord(wordSize int, alphabet string) string {
	ans := strings.Builder{}
	for i := 0; i < wordSize; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		ans.WriteByte(c)
	}
	return quote + ans.String() + quote
}

func GenerateRow(delimiter string, colsSize int, wordSize int, alphabet string) string {
	ans := []string{}
	for i := 0; i < colsSize; i++ {
		word := GenerateWord(wordSize, alphabet)
		ans = append(ans, word)
	}
	return strings.Join(ans, delimiter)
}

func GenerateRows(delimiter string, colsSize, rowsSize, wordSize int, alphabet string) []string {
	ans := []string{}
	for i := 0; i < rowsSize; i++ {
		row := GenerateRow(delimiter, colsSize, wordSize, alphabet)
		ans = append(ans, row)
	}
	return ans
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rows := GenerateRows(delimiter, colsSize, rowsSize, wordSize, alphabet)
	for i := 0; i < len(rows); i++ {
		fmt.Println(rows[i])
	}
}
