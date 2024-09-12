package main

import (
	"fmt"
	"time"
)

// 고루틴으로 처리할 작업 (1초 대기)
func task(count int) {
	fmt.Printf("Task %d 시작\n", id)
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Printf("Task %d 완료\n", id)
}

func withoutGoroutines() {
	fmt.Println("고루틴 없이 동기적으로 작업 시작")
	start := time.Now()

	// 동기적으로 5개의 작업을 수행
	for i := 1; i <= 5; i++ {
		task(i)
	}

	elapsed := time.Since(start)
	fmt.Printf("고루틴 없이 작업 완료. 소요 시간: %s\n\n", elapsed)
}

func withGoroutines() {
	fmt.Println("고루틴을 사용한 동시 작업 시작")
	start := time.Now()

	// 고루틴으로 5개의 작업을 동시 수행
	for i := 1; i <= 5; i++ {
		go task(i)
	}

	elapsed := time.Since(start)
	fmt.Printf("고루틴으로 작업 완료. 소요 시간: %s\n\n", elapsed)
}

func main() {
	// 동기적 작업
	withoutGoroutines()

	// 고루틴을 사용한 동시 작업
	withGoroutines()

	time.Sleep(5 * time.Second)
}
