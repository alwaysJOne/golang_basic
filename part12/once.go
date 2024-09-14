package main

import (
	"fmt"
	"sync"
)

func justOnce() {
	fmt.Println("한번만 실행하고싶은 함수")
}

func main() {
	//Once : 한 번만 실행(주로 초기화작업)
	var once sync.Once // Once 선언
	//once := new(sync.Once) //Once 객체 생성

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("고루틴 %d 실행\n", id)

			// 여러 고루틴이 동시에 호출해도 justOnce는 한 번만 실행됨
			once.Do(justOnce)
		}(i)
	}

	wg.Wait()
}
