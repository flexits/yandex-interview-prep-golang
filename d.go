/*
Дано целое число n. Требуется вывести все правильные скобочные последовательности длины 2 \* n, упорядоченные лексикографически.

Сложность O(C_n * 2n) по времени и O(n) по памяти, где C_n - число Каталана = (2n)! / (n! * (n+1)!)
*/
package main

import "fmt"

// Рекурсивно добавляем по одной скобке, до тех пор,
// пока не достигнем искомой длины 2 * count.
//
// Аргументы:
// current    текущая обрабатываемая строка;
// open       кол-во открывающих скобок;
// close      кол-во закрывающих скобок;
// count      требуемое кол-во скобок одного типа;
// resultPtr  указатель на слайс с результатами.
func generate(current string, open, close, count int, resultPtr *[]string) {
	// WARN: такое сравнение валидно только для ASCII
	if len(current) == 2*count {
		*resultPtr = append(*resultPtr, current)
		return
	}

	// Сначала ставим открывающую скобку и только потом закрывающую
	// для автоматической организации в лексикографическом порядке.
	if open < count {
		generate(current+"(", open+1, close, count, resultPtr)
	}
	if close < open {
		generate(current+")", open, close+1, count, resultPtr)
	}
}

func main() {
	count := 0
	if _, err := fmt.Scan(&count); err != nil {
		return
	}

	result := []string{}

	generate("", 0, 0, count, &result)

	for _, str := range result {
		fmt.Println(str)
	}
}

/*
// То же, но с меньшим кол-вом аргументов ф-ции:

func generate(count int) []string {
	result := []string{}

	var gen func(string, int, int)
	gen = func(current string, open, close int) {
		if len(current) == 2*count {
			result = append(result, current)
			return
		}

		if open < count {
			gen(current+"(", open+1, close)
		}
		if close < open {
			gen(current+")", open, close+1)
		}
	}

	gen("", 0, 0)
	return result
}

func main() {
	count := 3
	if _, err := fmt.Scan(&count); err != nil {
		return
	}

	result := generate(count)

	for _, str := range result {
		fmt.Println(str)
	}
}
*/