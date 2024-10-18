package main

import (
	"fmt"
	"time"
)

// 고루틴이 숫자를 생성하여 채널에 보냄
func getNumbers(ch chan<- int, id int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("고루틴 %d에서 %d 보냄\n", id, i)
		ch <- i                            // 채널에 숫자를 보냄
		time.Sleep(time.Millisecond * 500) // 0.5초 대기 (비동기성 강조)
	}
	close(ch) // 모든 데이터 전송 후 채널 닫기
}

func main() {
	ch1 := make(chan int) // 채널 1
	ch2 := make(chan int) // 채널 2

	// 두 개의 고루틴이 각각 다른 채널에 숫자를 생성
	go getNumbers(ch1, 1, 5) // 고루틴 1, 5개의 숫자 생성
	go getNumbers(ch2, 2, 3) // 고루틴 2, 3개의 숫자 생성

	// 채널에서 데이터를 받는 select 문
	for i := 0; i < 8; i++ { // 8번 반복 (총 8개의 값 수신)
		select {
		case num, ok := <-ch1:
			if ok {
				fmt.Println("채널 1에서 받은 값:", num)
			}
		case num, ok := <-ch2:
			if ok {
				fmt.Println("채널 2에서 받은 값:", num)
			}
		}
	}

	fmt.Println("모든 고루틴 완료 및 채널이 닫혔습니다.")
}
