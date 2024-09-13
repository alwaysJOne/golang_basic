package main

import (
	"fmt"
	"time"
)

func main() {
	//데이터를 보내는 고루틴이 값을 보내는 시점에 받는 고루틴이 준비되어 있어야 함
	//동기적 통신이므로 고루틴과 메인 고루틴이 정확히 데이터 전송 시점에서 만나야 작업이 진행
	//ch := make(chan int)

	// 버퍼 크기가 2인 채널 생성
	//버퍼를 사용하면 채널에 공간이 있으므로 대기하지않고 데이터를 즉시보냄 -> 비동기적 통신 지원
	ch := make(chan int, 2)

	go func() {
		fmt.Println("고루틴: 1을 보냅니다")
		ch <- 1 // 버퍼가 비어 있으므로 즉시 값을 보냄
		fmt.Println("고루틴: 1을 보냈습니다")

		fmt.Println("고루틴: 2를 보냅니다")
		ch <- 2 // 버퍼에 공간이 있으므로 즉시 값을 보냄
		fmt.Println("고루틴: 2를 보냈습니다")
	}()

	time.Sleep(2 * time.Second) // 1초 대기 (고루틴이 버퍼에 값을 보내는 동안)

	fmt.Println("메인 고루틴: 값을 받습니다")
	val1 := <-ch // 첫 번째 값 수신
	fmt.Printf("메인 고루틴: 받은 값 1: %d\n", val1)

	val2 := <-ch // 두 번째 값 수신
	fmt.Printf("메인 고루틴: 받은 값 2: %d\n", val2)
}
