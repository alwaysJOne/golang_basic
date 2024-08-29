package main

import (
	"fmt"
)

func main() {
	players := 5        // 총 플레이어 수 (사용자 + 컴퓨터)
	gameCount := 1      // 게임에서 현재 진행되는 숫자
	isGameOver := false // 게임 종료 플래그

	for !isGameOver {
		// 사용자 차례
		fmt.Printf("사용자 차례: ")
		var userInput string
		fmt.Scanln(&userInput)

		// 사용자 입력 검증
		isValid := true
		hasThreeSixNine := false

		// 숫자에 3, 6, 9가 포함되어 있는지 확인
		checkNum1 := gameCount % 10
		checkNum2 := gameCount / 10
		if checkNum1 == 3 || checkNum1 == 6 || checkNum1 == 9 ||
			checkNum2 == 3 || checkNum2 == 6 || checkNum2 == 9 {
			hasThreeSixNine = true
		}

		if hasThreeSixNine {
			if userInput != "x" {
				isValid = false
			}
		} else {
			expectedInput := fmt.Sprint(gameCount)
			if userInput != expectedInput {
				isValid = false
			}
		}

		if !isValid {
			fmt.Println("오답입니다. 게임 종료!")
			break
		}

		gameCount++

		// 컴퓨터 차례
		for i := 1; i < players; i++ {
			var computerInput string

			// 숫자에 3, 6, 9가 포함되어 있는지 확인
			checkNum1 := gameCount % 10
			checkNum2 := gameCount / 10

			if checkNum1 == 3 || checkNum1 == 6 || checkNum1 == 9 ||
				checkNum2 == 3 || checkNum2 == 6 || checkNum2 == 9 {
				computerInput = "x"
			} else {
				computerInput = fmt.Sprintf("%d", gameCount)
			}

			fmt.Printf("컴퓨터 %d 차례: %s\n", i, computerInput)
			gameCount++
		}
	}
}
