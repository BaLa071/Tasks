package main

import (
	"errors"
	"fmt"
)

func isNumber(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
func calculation(str string) (int, int) {
	var odd, even int = 0, 0
	flag := 1
	for i := len(str) - 1; i >= 0; i-- {
		if flag == 0 {
			odd += int(str[i] - '0')
		} else {
			temp := int(str[i]-'0') * 2
			rem := temp % 10
			temp /= 10
			even += rem + temp
			flag = 0
		}
	}
	return odd, even
}
func isCredit(str string) (string, error) {
	if len(str) >= 13 && len(str) <= 16 {
		if str[0] == '4' || str[0] == '5' || str[0] == '6' {
			if isNumber(str) {
				odd, even := calculation(str)
				if (odd+even)%10 == 0 {
					return "valid Credit card Number", nil
				} else {
					err := errors.New("Invalid Credit card Number2")
					return "", err
				}
			} else {
				err := errors.New("Invalid Credit card Number3")
				return "", err
			}
		} else {
			err := errors.New("Invalid Credit card Number4")
			return "", err
		}
	} else {
		err := errors.New("Invalid Credit card Number5")
		return "", err
	}
}
func main() {
	var str string
	fmt.Scan(&str)
	res, err := isCredit(str)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
