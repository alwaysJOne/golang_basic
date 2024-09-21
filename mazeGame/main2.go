package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"realCode/mazeGame/img"

	"github.com/hajimehoshi/ebiten/v2"
)

//캐릭터를 화면에 그려주기
//모든 이미지를 문자열로 변환하여 파일을 Go에 포함 할 수 있는 도구를 만들었습니다.
//이 도구 이름은 "file2dbyteslice" 빠른 로딩 속도를 제공합니다

//https://github.com/hajimehoshi/file2byteslice

//go build .\cmd\file2byteslice\main.go
//.\main.exe -input gopher.png -output gopher.go -package img -var Gopher_png

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{} //ebiten.DrawImageOptions 구조체의 포인터를 생성 -> 그릴때 이미지 옵션부여
	op.GeoM.Scale(1, 1)              // 가로와 세로 크기를 각각 1배
	op.GeoM.Translate(0, 0)          // 이미지를 화면의 (0, 0) 위치

	screen.DrawImage(gopher_img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	gopher_img *ebiten.Image //캐릭터이미지
)

const (
	screenWidth  = 500
	screenHeight = 500
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	// 이미지 객체를 생성 한다.
	gopher_, _, err := image.Decode(bytes.NewReader(img.Gopher_png))
	if err != nil {
		log.Fatal(err)
	}

	// go_char_img에 이미지 인터페이스를 연결
	gopher_img = ebiten.NewImageFromImage(gopher_)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
