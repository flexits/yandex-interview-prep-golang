/*
Дан упорядоченный по неубыванию массив целых 32-разрядных чисел. 
Требуется удалить из него все повторения.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()
	
	out, err := os.Create("output.txt")
	if err != nil {
		return
	}
	defer out.Close()

	in := bufio.NewScanner(file)
	in.Scan()
	line := in.Text()

	// первая строка файла содержит длину массива
	length, err := strconv.Atoi(line)
	if err != nil {
		return
	}
	if length <= 0 {
		return
	}
	
	// поскольку массив упорядочен по неубыванию согласно условию, 
	// задача сводится к тому, чтобы, проходя по входному массиву последовательно,
	// сравнить каждый новый элемент с последним уникальным элементом.
	
	// первый элемент входного массива помещаем в выходной без проверок
	in.Scan()
	line = in.Text()
	lastUnique, err := strconv.Atoi(line)
	if err != nil {
		return
	}
	//fmt.Println(lastUnique)
	fmt.Fprintln(out, lastUnique)

	// каждый последующий элемент сравниваем с последним добавленным в выходной файл
	for in.Scan() {
		line = in.Text()
		temp, err := strconv.Atoi(line)
		if err != nil {
			return
		}
		if temp > lastUnique {
			lastUnique = temp
			//fmt.Println(temp)
			fmt.Fprintln(out, temp)
		}
	}
}
