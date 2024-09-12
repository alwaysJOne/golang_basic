package main

import (
	"fmt"
	"log"
)

// 기본적으로 메서드 마다 리턴 타입 2개(리턴값, 에러)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// return 0, errors.New("0으로 나눌 수 없습니다.")
		return 0, fmt.Errorf("%.0f로 나눌 수 없습니다", b)
	}
	return a / b, nil
}

func main() {
	//Go에서는 함수가 오류가 발생할 수 있는 경우,
	//일반적으로 함수의 반환 값 중 하나로 오류를 반환함
	//오류가 발생할 수 있는 함수는 error 타입을 반환하며,
	//호출한 쪽에서 이 값을 검사하여 오류를 처리함

	var num1, num2 float64

	// 사용자로부터 두 개의 정수를 입력받음
	fmt.Print("두 개의 정수를 입력하세요 (예: 10 2): ")
	_, err := fmt.Scanf("%f %f", &num1, &num2)
	if err != nil {
		log.Fatal("입력 오류:", err)
	}

	result, err := divide(num1, num2)
	if err != nil {
		//에러 메시지를 출력한 후, 프로그램을 종료하는 함수
		log.Fatal(err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("========end===========")
}
