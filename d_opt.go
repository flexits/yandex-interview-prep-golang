/*
Дано целое число n. Требуется вывести все правильные скобочные последовательности длины 2 \* n, упорядоченные лексикографически.

Оптимизированная версия решения - без конкатенации строк.

Сложность O(C_n * 2n) по времени и O(n) по памяти, где C_n - число Каталана = (2n)! / (n! * (n+1)!)
*/
package main

import "fmt"

func generate(count int) []string {
	result := []string{}
	buffer := make([]byte, 2*count)

    // Рекурсивно добавляем по одной скобке, до тех пор,
    // пока не достигнем искомой длины 2 * count.
	// 
	// Аргументы:
    // idx    текущий индекс в буфере;
    // open   кол-во открывающих скобок;
    // close  кол-во закрывающих скобок.
	var gen func(idx, open, close int)
	gen = func(idx, open, close int) {
		if idx == 2*count {
			result = append(result, string(buffer))
			return
		}

		if open < count {
			buffer[idx] = '('
			gen(idx+1, open+1, close)
		}

		if open > close {
			buffer[idx] = ')'
			gen(idx+1, open, close+1)
		}
	}

	gen(0, 0, 0)
	return result
}

func main() {
	count := 0
	if _, err := fmt.Scan(&count); err != nil {
		return
	}

	result := generate(count)

	for _, str := range result {
		fmt.Println(str)
	}
}
