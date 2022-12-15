package main

import "fmt"

func main() {
	// 1～100のスライスを用意
	values := make([]int, 100)
	for i, _ := range values {
		values[i] = i + 1
	}
	fmt.Println(values)
	for _, v := range values {
		if v%3 == 0 && v%5 == 0 {
			fmt.Println("Fizz Buzz", v)
			continue
		}
		if v%3 == 0 {
			fmt.Println("Fizz", v)
			continue
		}
		if v%5 == 0 {
			fmt.Println("Buzz", v)
			continue
		}
		fmt.Println(v)
	}
}
