package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"realCode/mazeGame/img"

	"github.com/hajimehoshi/ebiten/v2"
)

//배경과 벽을 출력

//https://github.com/hajimehoshi/file2byteslice

//.\main.exe -input back.png -output back.go -package img -var Back_png
//.\main.exe -input wall.png -output wall.go -package img -var Wall_png

//2d화면을 그리기위한 2차원 슬라이스 설명

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	//맵을통해서  배경출력해주기
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			switch g_MapData[y][x] {
			case TypeWall:
				opWall := &ebiten.DrawImageOptions{}
				opWall.GeoM.Translate(float64(x+(x*50)), float64(y+(y*50)))
				screen.DrawImage(wallImg, opWall)
			case TypeBack:
				opBack := &ebiten.DrawImageOptions{}
				opBack.GeoM.Translate(float64(x+(x*50)), float64(y+(y*50)))
				screen.DrawImage(backImg, opBack)
			}
		}
	}

	op.GeoM.Translate(0, 0)

	screen.DrawImage(gopher_img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	gopher_img *ebiten.Image
	wallImg    *ebiten.Image //벽이미지
	backImg    *ebiten.Image //배경이미지

	//맵 데이터를 제작 하기 위해 맵 변수
	g_MapData [][]int
)

const (
	screenWidth  = 500
	screenHeight = 500

	TypeWall = 1 // 벽 타입 상수값
	TypeBack = 2 // 바닥 타입 상수값
)

func init() {
	g_MapData = [][]int{
		{2, 2, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 2, 2, 2, 2, 2, 1, 1, 2, 1},
		{1, 2, 2, 1, 1, 2, 2, 2, 2, 1},
		{1, 2, 2, 1, 1, 1, 1, 1, 2, 1},
		{1, 2, 2, 1, 1, 1, 1, 2, 1, 1},
		{1, 2, 2, 2, 2, 2, 2, 1, 2, 1},
		{1, 2, 2, 2, 2, 1, 2, 1, 2, 1},
		{1, 2, 2, 1, 2, 1, 2, 2, 2, 1},
		{1, 2, 2, 1, 2, 1, 1, 1, 2, 2},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	gopher_, _, err := image.Decode(bytes.NewReader(img.Gopher_png))
	if err != nil {
		log.Fatal(err)
	}

	//---------
	wall_, _, err := image.Decode(bytes.NewReader(img.Wall_png))
	if err != nil {
		log.Fatal(err)
	}
	//---------
	back_, _, err := image.Decode(bytes.NewReader(img.Back_png))
	if err != nil {
		log.Fatal(err)
	}

	gopher_img = ebiten.NewImageFromImage(gopher_)
	wallImg = ebiten.NewImageFromImage(wall_) //---------
	backImg = ebiten.NewImageFromImage(back_) //---------

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
