package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	var count int = 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'g' {
			if i < len(str)-1 && str[i+1] == 'o' {
				count++
			}
		}
	}
	fmt.Println(count)
}
