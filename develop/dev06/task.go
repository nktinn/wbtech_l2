package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fields    = flag.String("f", "0", "fields")
	delimit   = flag.String("d", "\t", "delimiter")
	separated = flag.Bool("s", true, "separated")
)

func main() {
	flag.Parse()

	// Читаем файл
	/*
		file, err := os.Open("develop/dev06/test.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)*/

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimit) {
			continue
		}
		columns := strings.Split(line, *delimit)
		for _, field := range strings.Split(*fields, ",") {
			if field == "0" {
				fmt.Println(line)
				break
			}
			if i, err := strconv.Atoi(field); err == nil && i < len(columns) {
				fmt.Print(columns[i], " ")
			}
		}
		fmt.Println()
	}

}
