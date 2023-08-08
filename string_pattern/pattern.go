package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	val := []byte(str)
	for i := 2; i < len(str); i += 4 {
		if i+1 > len(str)-1 {
			continue
		}
		val[i], val[i+1] = val[i+1], val[i]
	}
	fmt.Println(string(val))
}
