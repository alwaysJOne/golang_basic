package main

import (
	"fmt"
	"time"
)

func main() {
	//클로저 함수 : 함수와 그 함수가 접근할 수 있는 변수를 함께 기억하는 특별한 함수

	results := make([]int, 100)

	// 고루틴을 사용하여 각 인덱스의 제곱 값을 계산하는 클로저
	for i := 0; i < 100; i++ {
		go func() { // i를 인자로 전달하여 고유한 값을 사용
			time.Sleep(100 * time.Millisecond)
			results[i] = i // 고유한 i를 사용
		}() // 여기서 i를 인자로 전달
	}

	// 충분한 시간 대기
	time.Sleep(5 * time.Second)

	fmt.Println("결과:", results)
}
