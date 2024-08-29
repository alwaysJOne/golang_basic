package main

import "fmt"

func main() {
	/*
		if문
		특정 조건을 확인하여
		조건에 따라 다른 코드 블록을 실행할 필요가 있을 때 사용합니다.
	*/
	
	num := 10

	// 기본적인 if 문 사용 예제
	if num > 0 {
		fmt.Println("양수입니다.")
	}

	if num < 0 {
		fmt.Println("음수입니다.")
	}

	// if-else 문 사용 예제
	if num%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}

	// if-else if-else 문 사용 예제
	/*
		if num < 0 {
				fmt.Println("음수입니다.")
		} else {
			if num == 0 {
				fmt.Println("0 입니다.")
			} else {
				fmt.Println("양수입니다.")
			}
		}
	*/
	if num < 0 {
		fmt.Println("음수입니다.")
	} else if num == 0 {
		fmt.Println("0 입니다.")
	} else {
		fmt.Println("양수입니다.")
	}

	// if 문에서 초기화 구문 활용 예제
	if x := 5; x > 0 {
		fmt.Println("x는 양수입니다.")
		// 여기서만 x 변수 사용 가능
	}

	// 초기화 구문에서 선언한 변수는 if 문 외부에서 사용할 수 없음
	// fmt.Println(x) // 컴파일 에러: undefined: x

	//논리연산
	//조건식 && 조건식 (AND): 양쪽 모두가 참일 때 참을 반환합니다.
	//조건식 || 조건식 (OR): 어느 한 쪽이라도 참이면 참을 반환합니다.
	//!조건식 (NOT): 논리 상태를 반전시킵니다.
	//조건식에 자리에 보통 논리연산을 함께 사용하는 경우가 많다.
	num2 := 20

	// AND 연산자 (&&) 사용 예제
	if num2 > 0 && num2%2 == 0 {
		fmt.Println("양수이면서 짝수입니다.")
	}

	if num2 > 0 && num2%2 != 0 {
		fmt.Println("양수이면서 짝수입니다.")
	}

	// OR 연산자 (||) 사용 예제
	if num2 < 0 || num2%2 == 0 {
		fmt.Println("음수이거나 짝수입니다.")
	}

	if num2 < 0 || num2%2 != 0 {
		fmt.Println("음수이거나 짝수입니다.")
	}

	// NOT 연산자 (!) 사용 예제
	isPositive := num2 > 0
	if !isPositive {
		fmt.Println("음수입니다.")
	} else {
		fmt.Println("양수입니다.")
	}

	// 조건문에서 초기화 구문 활용과 AND 연산자 사용 예제
	if x := 5; x > 0 && x%2 == 0 {
		fmt.Println("x는 양수이면서 짝수입니다.")
		// 여기서만 x 변수 사용 가능
	}
}
