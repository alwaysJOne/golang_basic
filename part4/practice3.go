package main

import (
	"fmt"
)

func main() {
	// 사용자 이름 입력 받기
	var user1, user2 string
	fmt.Print("첫 번째 사용자의 이름을 입력하세요: ")
	fmt.Scanln(&user1)
	fmt.Print("두 번째 사용자의 이름을 입력하세요: ")
	fmt.Scanln(&user2)

	for {
		var userChoice1, userChoice2 string

		for {
			// 첫 번째 사용자에게 가위, 바위, 보 중 하나를 입력 받기
			fmt.Printf("\n%s님, 가위, 바위, 보 중 하나를 입력하세요: ", user1)
			fmt.Scanln(&userChoice1)
			if userChoice1 == "가위" || userChoice1 == "바위" || userChoice1 == "보" {
				break
			}

			fmt.Println("잘못입력하셨습니다. 가위, 바위, 보 중 하나를 입력하세요:")
		}

		for {
			// 두 번째 사용자에게 가위, 바위, 보 중 하나를 입력 받기
			fmt.Printf("\n%s님, 가위, 바위, 보 중 하나를 입력하세요: ", user2)
			fmt.Scanln(&userChoice2)
			if userChoice2 == "가위" || userChoice2 == "바위" || userChoice2 == "보" {
				break
			}

			fmt.Println("잘못입력하셨습니다. 가위, 바위, 보 중 하나를 입력하세요:")
		}

		// 사용자들의 선택 비교하여 결과 판정
		switch userChoice1 {
		case "가위":
			switch userChoice2 {
			case "가위":
				fmt.Println("무승부입니다. 다시 시도하세요.")
			case "바위":
				fmt.Printf("%s님이 이겼습니다!\n", user2)
				return
			case "보":
				fmt.Printf("%s님이 이겼습니다!\n", user1)
				return
			}
		case "바위":
			switch userChoice2 {
			case "가위":
				fmt.Printf("%s님이 이겼습니다!\n", user1)
				return
			case "바위":
				fmt.Println("무승부입니다. 다시 시도하세요.")
			case "보":
				fmt.Printf("%s님이 이겼습니다!\n", user2)
				return
			}
		case "보":
			switch userChoice2 {
			case "가위":
				fmt.Printf("%s님이 이겼습니다!\n", user2)
				return
			case "바위":
				fmt.Printf("%s님이 이겼습니다!\n", user1)
				return
			case "보":
				fmt.Println("무승부입니다. 다시 시도하세요.")
			}
		}
	}
}
