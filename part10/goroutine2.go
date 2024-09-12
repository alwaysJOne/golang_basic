package main

import (
	"fmt"
	"time"
)

// 고루틴으로 처리할 작업 (1초 대기)
func task(id string, count int) {
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Printf("%s : %d번 완료\n", id, count)
}

func withoutGoroutines() {
	// 동기적으로 5개의 작업을 수행
	for i := 1; i <= 5; i++ {
		task("test1", i)
	}
}

func withGoroutines() {
	// 고루틴으로 5개의 작업을 동시 수행
	for i := 1; i <= 5; i++ {
		go task("test2", i)
	}
}

func main() {
	// 동기적 작업
	withoutGoroutines()

	// 고루틴을 사용한 동시 작업
	withGoroutines()

	time.Sleep(5 * time.Second)
}
