/*
Даны две строки строчных латинских символов: строка J и строка S. 
Нужно определить, какое количество символов из S входит в J.
*/
package main

import "fmt"

func main() {
	var s, j string
	fmt.Scanf("%s", &j)
	fmt.Scanf("%s", &s)

	// составить таблицу искомых символов из строки J
	chars := make(map[rune]struct{})
	for _, sym := range j {
		chars[sym] = struct{}{}
	}

	// проверить каждый элемент строки S:
	// если он входит в число искомых, увеличить счётчик совпадений
	result := 0
	for _, sym := range s {
		_, ok := chars[sym]
		if ok {
			result++
		}
	}

	fmt.Println(result)
}