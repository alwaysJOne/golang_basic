package main

import (
	"fmt"
)

func main() {
	//recover: panic에 의해 프로그램이 종료되는 것을 막고, 프로그램을 복구할 때 사용
	//오직 defer 함수 내에서만 호출될 수 있음
	defer func() {
		//recover() : 패닉이 발생하지 않았다면 nil을 반환함
		if r := recover(); r != nil {
			fmt.Println("복구 완료! 패닉 메시지:", r)
		}
	}()

	//Go 언어에는 panic과 recover키워드를 통해 에러처리를 지원
	//panic이 호출되면 현재 함수가 즉시 중지되며, defer 함수를 호출하고 자기자신을 호출한 곳으로 리턴
	//논리에러 발생 처리 가능

	fmt.Println("프로그램 시작")
	panic("예상가능한 논리에러 발생! 프로그램중단")
	fmt.Println("프로그램 종료") //실행 불가
}
