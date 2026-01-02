/*
Требуется найти самую длинную последовательность единиц и вывести её длину.
Желательно получить решение, работающее за линейное время и при этом 
проходящее по входному массиву только один раз.
*/
package main

import "fmt"

func main() {
	length, temp, result, counter := 0, 0, 0, 0
	fmt.Scanf("%d", &length)
	
	// для каждой последовательности единиц 
	// подсчитываем длину и обновляем наибольшее
	for range length {
		fmt.Scanf("%d", &temp)
		if temp == 1 {
			counter++
			if result < counter {
				result = counter
			}
		} else {
			counter = 0
		}
	}

	fmt.Println(result)
}