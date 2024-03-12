package main

import "fmt"

func main() {
	fmt.Println(rescue(5, 5, []int{2, 5, 10, 12, 15}))       // 2
	fmt.Println(rescue(6, 10, []int{1, 11, 30, 34, 35, 37})) // 4
}
