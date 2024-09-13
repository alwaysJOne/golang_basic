package main

import (
	"fmt"
)

func getNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // 채널에 숫자를 보냄
	}
	//Close : 더이상 채널에 데이터를 보내지 않을 때 사용
	//		  닫힌 채널에 값을 수신할 수 있지만 전송은 불가, 전송시 패닉(예외) 발생
	close(ch) // 모든 데이터 전송 후 채널 닫기
}

func main() {
	ch := make(chan int, 5)

	go getNumbers(ch) // 고루틴에서 숫자 생성

	// range를 사용하여 채널에서 값을 받음
	//Range : 채널을 닫을 때 까지 반복해서 값을 꺼냄, 채널을 닫지않으면 데드락(무한 대기 상태)발생
	for num := range ch {
		fmt.Println("받은 값:", num)
	}

	fmt.Println("채널이 닫혔습니다.")
}
