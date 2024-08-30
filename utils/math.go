// utils/math.go
package utils

import "fmt"

// Add 두 정수를 더하는 함수
func Add(a int, b int) int {
	return a + b
}

// Add 두 정수를 더해서 출력하는 메소드
func PrintAdd(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

//패키지에 변수, 함수, 타입등 내부 자원의 이름이 소문자일 때는 외부에서 접근이 불가함
func printMinus(a int, b int) int {
	return a - b
}
