package main

import "fmt"

func main() {
	// 부호 있는 정수형
	var num1 int = 10
	num1 = -10
	num1 = 0
	fmt.Println("정수 :", num1)

	// 부호 없는 정수형
	var num2 uint = 20
	// num2 = -20
	fmt.Println("부호 없는 정수 :", num2)

	// 다양한 정수형
	var num3 int8 = 30
	var num4 int16 = 40
	var num5 int32 = 50
	var num6 int64 = 60
	fmt.Println("다양한 정수형 :", num3, num4, num5, num6)

	// 정수형 간의 형변환 예제
	var result int32
	result = num5 + 20
	result = int32(num3) + int32(num4)
	fmt.Println("정수형 간의 형변환 :", result)

	//실수형
	var f1 float32 = 3.14
	var f2 float64 = 3.141592653589793
	fmt.Println("실수형 :", f1, f2)
	fmt.Println("실수형 :", int(f1), int(f2))

	// 복소수
	var com1 complex64 = 1 + 2i
	var com2 complex128 = 3 + 4i
	fmt.Println("복소수 :", com1, com2)

	// 문자열
	var str1 string = "Hello, Go!"
	fmt.Println("문자열 :", str1)

	// 논리형
	var b1 bool = true
	var b2 bool = false
	fmt.Println("부울 값 :", b1, b2)
	b3 := (num1 < 10)
	fmt.Println("비교연산 결과 : ", b3)

	// byte
	var byte1 byte = 'a'
	var byte2 byte = 65
	fmt.Printf("바이트 : %c, %c\n", byte1, byte2)

	// byte 타입에서 int로의 형변환 예제
	var byteToInt int
	byteToInt = int(byte1)
	fmt.Println("바이트를 int로 형변환 :", byteToInt)
}
