package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const (
	SheetName      = "Лист1"
	OutputFileName = "stations.csv"
)

func main() {
	modifyRow := func(row []string) ([]string, error) {
		if len(row) != 5 {
			return []string{}, fmt.Errorf("error! Wrong lenght of the row: %s (%d)", row, len(row))
		}
		ans := make([]string, 5)
		ans[0] = row[0]
		ans[1] = row[2]
		ans[2] = row[1]
		ans[3] = row[4]
		return ans, nil
	}

	modifyRows := func(rows [][]string) ([][]string, error) {
		ans := make([][]string, len(rows)-1)
		for i := 1; i < len(rows); i += 1 {
			if len(rows[i]) == 0 {
				continue
			}
			modifiedRow, err := modifyRow(rows[i])
			if err != nil {
				return nil, err
			}
			ans[i-1] = modifiedRow
		}
		return ans, nil
	}

	getAllRows := func() ([][]string, error) {
		file, err := excelize.OpenFile("stations.xlsx")
		if err != nil {
			return nil, err
		}
		defer file.Close()

		rows, err := file.GetRows(SheetName)
		if err != nil {
			return nil, err
		}

		return modifyRows(rows)
	}

	getCSVOutput := func(rows [][]string) ([]byte, error) {
		buf := &bytes.Buffer{}

		writer := csv.NewWriter(buf)
		defer writer.Flush()
		writer.Comma = ';'

		err := writer.WriteAll(rows)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	saveToFileWithCP866Encoding := func(filename string, text []byte) error {
		fd, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer fd.Sync()
		defer fd.Close()

		bufferedWriter := bufio.NewWriter(fd)
		defer bufferedWriter.Flush()

		w := transform.NewWriter(bufferedWriter, charmap.CodePage866.NewEncoder())
		defer w.Close()

		n, err := w.Write(text)
		if err != nil || n != len(text) {
			return err
		}

		return nil
	}

	rows, err := getAllRows()
	if err != nil {
		panic(err)
	}

	text, err := getCSVOutput(rows)
	if err != nil {
		panic(err)
	}

	if err := saveToFileWithCP866Encoding(OutputFileName, text); err != nil {
		panic(err)
	}
}
