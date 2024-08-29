package main

import "fmt"

func main() {
	/*
		for문
		반복적인 작업을 수행할 때
	*/
	// 기본적인 for 문 사용 예제
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	// 초기화 구문과 증감식부분을 생략하고 작성
	j := 0
	for j < 5 {
		fmt.Print(j, " ")
		j++
	}

	fmt.Println()
	// 무한 루프 (조건식 생략) -> 보통 정해진 횟수가 없을 때 사용
	k := 0
	for {
		fmt.Print(k, " ")
		k++
		if k >= 5 {
			break // 무한 루프를 빠져나가는 조건
		}
	}

	fmt.Println()
	// continue 문을 사용한 예제
	for m := 1; m <= 10; m++ {
		if m == 2 {
			continue // m이 2일 때 현재 반복을 중지하고 다음 반복을 시작
		}
		fmt.Print(m, " ")
	}

}
