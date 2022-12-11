package main

import (
	"bufio"
	"fmt"
	"strings"
)

var data = `
45083	Feature	Testing	Normal	RAIL: PG: Включить последние таблицы в сборку через pg	2022-12	07.12.2022 19:35
`

func TaskFromLine(line string) (string, string) {
	items := strings.Split(line, " ")
	task := "#" + items[0]
	result := append(append([]string{}, task), items[4:len(items) - 3]...)
	return task, strings.Join(result, " ")
}

func main() {
	reader := strings.NewReader(strings.Trim(data, " \n\t"))
	scanner := bufio.NewScanner(reader)
	tasks := []string{}
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\t", " ")
		task, name := TaskFromLine(line)
		tasks = append(tasks, task)
		fmt.Println(name)
	}
	fmt.Println(tasks)
}