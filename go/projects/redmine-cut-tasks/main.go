package main

import (
	"bufio"
	"fmt"
	"strings"
)

var data = `
45239	Bug	Resolved	Normal	RAIL: PG: Исправить ошибку на авто транзакцию для railhist при обработке ошибок	2022-12	21.12.2022 19:33
45236	Feature	Resolved	Normal	RAIL: PG: Реализовать механизм авто транзакции в библиотеке rail curs	2022-12	21.12.2022 19:33
45056	Feature	Resolved	Normal	RAIL: Исправить латиницу в Заявке на приобретение билета	2022-12	21.12.2022 19:33
44987	Feature	Resolved	Normal	RAIL: Оповещение при возврате нескольких билетов из заказа	2022-12	21.12.2022 19:33
44881	Bug	Resolved	Normal	RAIL: НЕМО: Международный класс обслуживания	2022-12	21.12.2022 19:33
42337	Feature	Resolved	Normal	RAIL: PDF: Придумать механизм более быстрой загрузки файлов для агентов	2022-12	21.12.2022 19:33
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