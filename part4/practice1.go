package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// 첫 번째 숫자 입력
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	fmt.Scanln(&num1)

	// 두 번째 숫자 입력
	fmt.Print("두 번째 숫자를 입력하세요: ")
	fmt.Scanln(&num2)

	// 연산자 입력 (+, -, *, /)
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	fmt.Scanln(&operator)

	// 연산자에 따른 연산 수행
	switch operator {
	case "+":
		fmt.Printf("결과 : %.2f\n", num1+num2)
	case "-":
		fmt.Printf("결과 : %.2f\n", num1-num2)
	case "*":
		fmt.Printf("결과 : %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("0으로 나눌 수 없습니다.")
		} else {
			fmt.Printf("결과 : %.2f\n", num1/num2)
		}
	default:
		fmt.Println("올바른 연산자를 입력하세요 (+, -, *, /)")
	}
}
