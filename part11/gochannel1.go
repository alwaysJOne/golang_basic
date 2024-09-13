package main

import (
	"fmt"
	"time"
)

func task(id string, ch chan string) {
	fmt.Printf("Worker %s: 작업 시작\n", id)
	time.Sleep(time.Second) // 작업을 1초 동안 수행
	fmt.Printf("Worker %s: 작업 완료\n", id)

	//채널 <- 데이터 : 채널에 데이터 전송
	ch <- id // 작업 완료 후, 채널에 id 전송
}

func main() {
	//채널
	//고루틴 간의 통신을 위한 도구로, 데이터를 안전하게 주고받을 수 있는 방법

	//데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고받을 수 있음)
	ch := make(chan string)

	go task("t1", ch) // 3개의 고루틴 실행
	go task("t2", ch) // 3개의 고루틴 실행
	go task("t3", ch) // 3개의 고루틴 실행

	// 3개의 고루틴으로부터 작업 완료 신호 받기
	id := <-ch // 채널에서 id를 수신
	fmt.Printf("task %d의 작업 완료 신호 수신\n", id)

	id = <-ch // 채널에서 id를 수신
	fmt.Printf("task %d의 작업 완료 신호 수신\n", id)

	id = <-ch // 채널에서 id를 수신
	fmt.Printf("task %d의 작업 완료 신호 수신\n", id)

	fmt.Println("모든 작업 완료")
}
