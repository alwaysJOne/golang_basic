package main

import "fmt"

func main() {
	/*
		switch문
		여러 조건을 확인하여 코드를 실행해야 할 때 또는
		일치하는 값에 따른 코드를 실행해야할 때 사용
	*/

	num := 10

	// 기본적인 switch 문 사용 예제
	switch num {
	case 1:
		fmt.Println("숫자는 1입니다.")
	case 2:
		fmt.Println("숫자는 2입니다.")
	case 3:
		fmt.Println("숫자는 3입니다.")
	default:
		fmt.Println("1~3사이의 숫작가 아닙니다.")
	}

	// switch 문에서 초기화 구문 활용 예제
	switch lang := "go"; lang {
	case "java":
		fmt.Println("최고의 언어 java")
	case "python":
		fmt.Println("최고의 언어 python")
	case "go":
		fmt.Println("최고의 언어 go")
	case "sql":
		fmt.Println("최고의 언어 sql")
	}

	// switch 문에서 값전달 없이 조건식을 통해서 사용하는 방법 (switch true)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A 학점")
	case score >= 80:
		fmt.Println("B 학점")
	case score >= 70:
		fmt.Println("C 학점")
	case score >= 60:
		fmt.Println("D 학점")
	default:
		fmt.Println("F 학점")
	}

	//여러값에 동일한 결과를 실행하는 방법
	switch num3 := 5; num3 {
	case 1, 2:
		fmt.Println("x는 1, 2")
	case 3, 4:
		fmt.Println("x는 3, 4")
	case 5:
		fmt.Println("x는 5")
		fallthrough
	case 6:
		fmt.Println("x는 6")
	}
}
