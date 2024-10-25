package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var sum int
var Global = 5

func main() {
	// WordCount("Mary had a little little lamb lamb lamb")
	// maps()
	// data := []int{33, 10, 55, 26, 7, 1, 59, 19}
	// fmt.Println("Unsorted:", data)
	// fmt.Println("Sorted:", quickSort(data))
	// more_slices()
	// fmt.Println(products([]string{"bread", "ham", "cheese", "cucumbers"}))
	// fmt.Println(two_sum([]int{1, 2, 3, 4, 5}, 5))
	// fmt.Println(maps([]string{"cat", "dog", "bird", "dog", "parrot", "cat"}))
	// fmt.Println(Fib(5))
	// iterator := Generate(0)
	// iterator()
	// iterator()
	// iterator()
	// iterator()
	// iterator()
	// fmt.Println(add(5))
	// fmt.Println(add(10))
	// fmt.Println(add(20))
	// ar, ok := Area(myFigure)
	// fmt.Printf("%s", unintuitive())
	// VaryLongTimeFunction()
	fmt.Println(Global)
	deferic()
	fmt.Println(Global)
}

func credit() {
	var x, p, y, year int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	x = x * 100

	for x < (y * 100) {
		x = x + (x * p / 100)
		year += 1
	}
	print(year)
}

func printf_practice() {
	var i float64
	fmt.Scan(&i)
	if i > 0 {
		if i > 10000 {
			fmt.Printf("%e", i)
		} else {
			i *= i
			fmt.Printf("%.4f", i)
		}
	} else {
		fmt.Printf("число %2.2f не подходит", i)
	}
}

func for_fun() {
	ch := "emoji"
	fmt.Println([]rune(ch))
}

func arrays() {
	// var new_array [3]int = [3]int{1, 2, 3}
	arr := [5]int{1, 2, 3, 4, 5}
	// new_array_3 := [3]int{2: 12, 1: 15}
	for idx, elem := range arr {
		fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
	}

}

func slices() {
	// arr := [5]int{1, 2, 3, 4, 5}
	slice := make([]int, 3)

	fmt.Println(slice)
}

func binary_search(arr [10]int, searched int) (item int) {
	var (
		high int = len(arr) - 1
		low  int = 0
	)

	for low <= high {
		mid := (high + low) / 2

		if arr[mid] == searched {
			item = mid
			break
		} else if arr[mid] > searched {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return item
}

func Sqrt(num float64) float64 {
	guess := 1.0
	i := 0
	for {
		i++
		guess -= (guess*guess - num) / (2 * guess)
		fmt.Println("Attemps #", i, " ", guess)
		if guess == math.Sqrt(num) {
			break
		}
	}
	return guess
}

func switch_exercise() {
	var size string
	fmt.Println("Enter your size :)")
	fmt.Scan(&size)

	switch size {
	case "S":
		fmt.Println("Wow! You are so thick!")
	case "M":
		fmt.Println("Its good!")
	default:
		fmt.Println("YOU ARE FAT OGR!!!")
	}
}

func slices_2() {
	slice := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	for i, v := range slice {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pow := make([][]uint8, dy)

	for i := range pow {
		pow[i] = make([]uint8, dx)
		for j := range pow[i] {
			pow[i][j] = uint8((i ^ j) * (i ^ j))
		}
	}
	return pow
}

func WordCount(s string) {
	words := strings.Fields(s)
	hashMap := map[string]int{}

	for _, elem := range words {
		if hashMap[elem] == 0 {
			hashMap[elem] = 1
		} else {
			hashMap[elem] += 1
		}
	}
	fmt.Println(hashMap)
}

func struct_2() {
	type MyStruct struct {
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func more_slices() {
	dim := 100
	s := make([]int, 0, dim)

	for i := 0; i < dim; i++ {
		s = append(s, i+1)
	}

	fmt.Println(s)

	s = append(s[:10], s[dim-10:]...)

	fmt.Println(s)

	dim = len(s)

	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}

	fmt.Println(s)
}

func products(order []string) (int, []string) {
	var (
		total_cost int
		delicacies []string
		products   = map[string]int{
			"bread":     50,
			"milk":      100,
			"butter":    200,
			"sausage":   500,
			"salt":      20,
			"cucumbers": 200,
			"cheese":    600,
			"ham":       700,
			"pork":      900,
			"tomatoes":  250,
			"fish":      300,
			"jamon":     1500,
		}
	)

	for _, item := range order {
		if products[item] != 0 {
			total_cost += products[item]
		}
	}

	for key := range products {
		if products[key] > 500 {
			delicacies = append(delicacies, key)
		}
	}

	return total_cost, delicacies
}

func two_sum(A []int, k int) []int {
	var (
		answer = make([]int, 2)
		elems  = make(map[int]int)
		diff   int
	)

	for idx := range A {
		elems[A[idx]] = idx
	}
	fmt.Println(elems)

	for key := range elems {
		diff = k - key
		if elems[diff] != 0 {
			answer[0], answer[1] = elems[diff], elems[key]
		}
	}

	return answer
}

func maps(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)
		}
		inputSet[v] = struct{}{}
	}
	fmt.Println(inputSet)
	return output
}

func Fib(n int) int {
	switch {
	case n <= 1: // терминальная ветка
		return n
	default: // рекурсивная ветка
		return Fib(n-1) + Fib(n-2)
	}
}

func Generate(seed int) func() {
	return func() {
		fmt.Println(seed) // замыкание получает внешнюю переменную seed
		seed += 2         // переменная модифицируется
	}

}

func add(x int) int {
	sum += x
	return sum
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if predicate(filename) {
				fmt.Println(filename)
			}

			if f.IsDir() {
				walk(filename)
			}
		}
	}
}

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}

func VaryLongTimeFunction() {
	sum := 0
	defer metricTime(time.Now())
	for i := 1; i < 1000000000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func deferic() {
	defer func(checkout int) {
		Global = checkout
	}(Global)
	Global = 42
	fmt.Println(Global)
}
