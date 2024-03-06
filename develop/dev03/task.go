package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	columnFlag = flag.Int("k", 1, "указание колонки для сортировки")
	numFlag    = flag.Bool("n", false, "сортировать по числовому значению")
	revFlag    = flag.Bool("r", false, "сортировать в обратном порядке")
	uniqFlag   = flag.Bool("u", false, "не выводить повторяющиеся строки")
)

func main() {
	flag.Parse()

	lines := readLines()
	lines = sortLines(lines)
	printLines(lines)
}

func readLines() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func sortLines(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		fields1 := strings.Fields(lines[i])
		fields2 := strings.Fields(lines[j])

		col1 := getField(fields1, *columnFlag)
		col2 := getField(fields2, *columnFlag)

		if *numFlag {
			num1, err1 := strconv.Atoi(col1)
			num2, err2 := strconv.Atoi(col2)
			if err1 == nil && err2 == nil {
				return num1 < num2
			}
		}

		return col1 < col2
	})

	if *revFlag {
		reverse(lines)
	}

	if *uniqFlag {
		lines = unique(lines)
	}

	return lines
}

// Получить элемент по номеру колонки

func getField(fields []string, col int) string {
	if col > len(fields) {
		return ""
	}
	return fields[col-1]
}

// Сортировка в обратном порядке

func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// Удаление повторяющихся строк

func unique(lines []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}
	return result
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
