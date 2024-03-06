package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	anagrams := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		runes := []rune(word)
		slices.Sort(runes)
		sortedWord := string(runes)

		if !slices.Contains(anagrams[sortedWord], word) {
			anagrams[sortedWord] = append(anagrams[sortedWord], word)
		}
	}

	// Формируем мапу множеств анаграмм
	for _, words := range anagrams {
		if len(words) == 1 {
			continue
		}

		sort.Strings(words)
		key := words[0]
		anagramSets[key] = words
	}

	return anagramSets
}

func main() {
	words := []string{"Пятак", "пятка", "тяпка", "птк", "листок", "слиток", "столиК", "листок"}

	fmt.Println(findAnagramSets(words))
}
