package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("파일열람실패 : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}

	defer f.Close()
}

func main() {
	// 보통패닉은 db연결실패, 파일 경로 불일치등 프록그램진행이 어려운 심각한 경우 사용
	fileOpen("log.txt")
	fmt.Println("로그 파일을 성공적으로 열었습니다.")
}
