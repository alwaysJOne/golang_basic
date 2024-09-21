package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math"
	"realCode/mazeGame/img"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//벽으로이동할 수 없고 다음끝을 찾아 이동한 후 게임종료되게 수정

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	for _, key := range g.keys {
		newX, newY := g_Character.GetPosX(), g_Character.GetPosY() // 캐릭터 좌표 초기화
		mapX, mapY := newX/50, newY/50                             //x좌표 y좌표로 변환
		switch key.String() {
		case "ArrowLeft":
			newX -= Speed                              // left 키 입력시 속도만큼 x좌표 차감
			mapX = int(math.Floor(float64(newX) / 50)) //맵의 x
		case "ArrowRight":
			newX += Speed                             // right 키 입력시 속도만큼 x좌표 증가
			mapX = int(math.Ceil(float64(newX) / 50)) //맵의 x
		case "ArrowUp":
			newY -= Speed                              // up 키 입력시 속도만큼 y좌표 차감
			mapY = int(math.Floor(float64(newY) / 50)) //맵의 y
		case "ArrowDown":
			newY += Speed                             // down 키 입력시 속도만큼 y좌표 증가
			mapY = int(math.Ceil(float64(newY) / 50)) //맵의 y
		}

		// 목표 위치에 도달했는지 확인 (8, 9)
		if mapY == 8 && mapX == 9 {
			return fmt.Errorf("game exited") // 게임 종료
		}

		// 새로운 위치가 벽인지 확인
		if newX >= 0 && newX < screenWidth && newY >= 0 && newY < screenHeight {

			if g_MapData[mapY][mapX] != TypeWall {

				g_Character.SetPosInfo(newX, newY) // 이동이 가능하면 캐릭터 위치 변경
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)

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

	op.GeoM.Translate(float64(g_Character.GetPosX()), float64(g_Character.GetPosY()))

	screen.DrawImage(gopher_img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

type ObjectPosInfo struct {
	x int
	y int
}

func (o *ObjectPosInfo) GetPosX() int { return o.x }
func (o *ObjectPosInfo) GetPosY() int { return o.y }
func (o *ObjectPosInfo) SetPosInfo(x int, y int) {
	o.x = x
	o.y = y
}

func (o *ObjectPosInfo) AddPosX(x int) { o.x += x }
func (o *ObjectPosInfo) AddPosY(y int) { o.y += y }

type Character struct {
	ObjectPosInfo
}

var (
	gopher_img *ebiten.Image
	wallImg    *ebiten.Image
	backImg    *ebiten.Image

	g_Character Character
	g_MapData   [][]int
)

const (
	screenWidth  = 500
	screenHeight = 500

	TypeWall = 1
	TypeBack = 2

	Speed                 = 2
	DefaultStartPositionX = 0
	DefaultStartPositionY = 0
)

func init() {
	g_Character.SetPosInfo(DefaultStartPositionX, DefaultStartPositionY)

	// 맵에 벽과 바닥 배치
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

	wall_, _, err := image.Decode(bytes.NewReader(img.Wall_png))
	if err != nil {
		log.Fatal(err)
	}

	back_, _, err := image.Decode(bytes.NewReader(img.Back_png))
	if err != nil {
		log.Fatal(err)
	}

	gopher_img = ebiten.NewImageFromImage(gopher_)
	wallImg = ebiten.NewImageFromImage(wall_)
	backImg = ebiten.NewImageFromImage(back_)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
