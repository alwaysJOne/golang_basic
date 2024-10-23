package main

import (
	"bytes"
	"fmt"
	"goplus/mazeGame/img"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

//화면에 fPS를 출력해보자
// 1초당 보여주는 화면의 장수
//go get golang.org/x/image => 폰트를 사용하여 텍스트를 그래픽으로 렌더링할 수 있는 기능

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(100, 200)

	screen.DrawImage(gopher_img, op)

	msg := fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()) // msg 변수에 FPS 문자열을 저장

	d := &font.Drawer{
		Dst:  screen,                        //`Dst`는 텍스트를 그릴 대상 이미지
		Src:  image.NewUniform(color.White), //`Src`는 텍스트의 색상을 정의
		Face: normalFont,                    // `Face`는 텍스트를 렌더링할 때 사용할 폰트
		Dot:  fixed.Point26_6{X: fixed.I(10), Y: fixed.I(30)},
		//`Dot`은 텍스트를 그리기 시작할 위치(고정 소수점 형식)
	}
	//정의된 위치와 폰트로 주어진 문자열을 그리는 함수
	d.DrawString(msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	gopher_img *ebiten.Image
	normalFont font.Face // 추가
)

const (
	screenWidth  = 500
	screenHeight = 500
)

func init() {
	//본 제공 트루타입 폰트를 파싱하여 tt라는 변수에 저장
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	//화면에 텍스트를 렌더링할 때의 해상도
	const dpi = 96
	//주어진 폰트를 기반으로 새로운 `Face`(글꼴의 스타일)를 만듭니다
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull, //텍스트를 더 선명하게 렌더링
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	gopher_, _, err := image.Decode(bytes.NewReader(img.Gopher_png))
	if err != nil {
		log.Fatal(err)
	}

	gopher_img = ebiten.NewImageFromImage(gopher_)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
