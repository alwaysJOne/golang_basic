package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// 공유 자원
	counter int
	// 뮤텍스
	mutex sync.Mutex
	// 원스
	once sync.Once
)

func initialize() {
	// 초기화 작업 (단 한 번만 실행됨)
	counter = 0
	fmt.Println("공유 자원 초기화 완료")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 작업이 완료되면 Done 호출

	// 초기화 작업을 한 번만 수행
	once.Do(initialize)

	mutex.Lock()         // 뮤텍스 잠금
	defer mutex.Unlock() // 함수 종료 시 잠금 해제

	// 공유 자원을 안전하게 수정
	counter++
	fmt.Printf("작업자 %d: counter = %d\n", id, counter)
	time.Sleep(time.Millisecond * 100) // 작업을 시뮬레이션
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)         // 각 고루틴을 실행하기 전에 WaitGroup에 1 추가
		go worker(i, &wg) // 고루틴 시작
	}

	wg.Wait() // 모든 고루틴이 완료될 때까지 대기
	fmt.Println("모든 작업 완료. 최종 counter 값:", counter)
}
