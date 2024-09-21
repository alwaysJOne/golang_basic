package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//ebitengine을 사용해서 화면에 게임 화면을 띄워 보기
//go get "github.com/hajimehoshi/ebiten/v2"으로 패키지 다운로드

type Game struct { // 게임을 만들기 위해서 꼭 만들어야 하는 필수 객체 입니다.
}

//	Update, Draw는 매프레임마다 실행됨
//
// 이 업데이트 함수는 에빗엔진 내부에서 자동으로 호출합니다.
func (g *Game) Update() error {
	return nil
}

// 이 함수를 화면에 그림을 그릴때 호출되는 함수 입니다.
func (g *Game) Draw(screen *ebiten.Image) {
	// 앞으로 여기에 그림을 그릴 예정
}

// 이 함수는 에빗 엔진 라이브러리 내부에 화면사이즈를 이렇게 할 거라고 알려주는 역활을 합니다.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight // 우리가만든 화면 사이즈를 엔진에 알려줌.
}

const (
	screenWidth  = 500 // 가로 사이즈
	screenHeight = 500 // 세로 사이즈
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight) // 화면사이즈 설정하기
	// ebiten.RunGame()이라는 함수를 이용해서 우리가 만든 Game 객체 메모리를 생성해서
	// 에빗엔진에게 던져 줍니다. 그럼 에빗 엔진은 Game 객체의 설정내용을 분석해서 동작을 시작 합니다.
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
