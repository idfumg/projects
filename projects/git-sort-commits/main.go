package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

const (
	GIT_SHOW_DATE_COMMAND  = "git show -s --format=%ci "
	GIT_SHOW_TITLE_COMMAND = "git show -s --format=%s "
	DATE_LAYOUT            = `2006-01-02 15:04:05 +0300`
)

type Record struct {
	Title     string
	Hash      string
	Timestamp time.Time
}

type Records []Record

func execShell(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func readLines(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(string(content), "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], " \n\t")
	}

	ans := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		ans = append(ans, line)
	}

	return ans, nil
}

func parseRecord(hash string) (Record, error) {
	date, stderr, err := execShell(GIT_SHOW_DATE_COMMAND + hash)
	if err != nil {
		return Record{}, fmt.Errorf("%s (%s)", strings.Trim(stderr, " \n\t"), err)
	}

	date = strings.Trim(date, " \n")

	t, err := time.Parse(DATE_LAYOUT, date)
	if err != nil {
		return Record{}, err
	}

	title, stderr, err := execShell(GIT_SHOW_TITLE_COMMAND + hash)
	if err != nil {
		return Record{}, fmt.Errorf("%s (%s)", strings.Trim(stderr, " \n\t"), err)
	}

	title = strings.Trim(title, " \n")

	return Record{
		Title:     title,
		Hash:      hash,
		Timestamp: t,
	}, nil
}

func parseRecords(filename string) (Records, error) {
	lines, err := readLines(filename)
	if err != nil {
		return Records{}, err
	}

	records := Records{}
	for _, line := range lines {
		record, err := parseRecord(line)
		if err != nil {
			return Records{}, err
		}

		records = append(records, record)
	}
	return records, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Args[0], "file_with_commit_hashes")
		os.Exit(1)
	}
	filename := os.Args[1]

	records, err := parseRecords(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Timestamp.UTC().Nanosecond() < records[j].Timestamp.UTC().Nanosecond()
	})

	for _, record := range records {
		fmt.Println(record.Title)
		fmt.Println(record.Hash)
		fmt.Println()
	}
}
