package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)
	var pascal = make([][]int, input)
	for i := 0; i < input; i++ {
		pascal[i] = make([]int, i+1)
	}
	for i := 0; i < input; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				pascal[i][j] = 1
			} else {
				pascal[i][j] = pascal[i-1][j] + pascal[i-1][j-1]
			}
		}
	}
	temp := input + 1
	for i := 0; i < input; i++ {
		temp--
		for j := 0; j < temp; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < len(pascal[i]); j++ {
			fmt.Print(pascal[i][j], " ")
		}
		fmt.Println()
	}

}
