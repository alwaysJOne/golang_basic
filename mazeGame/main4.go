package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"realCode/mazeGame/img"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil" //키 입력 패키지
)

//캐릭터움직이기

type Game struct {
	keys []ebiten.Key // <---- 추가
}

func (g *Game) Update() error {
	//g.keys에는 이전에 눌렸던 키목록이 존재
	// 주기적으로 키 입력을 감지해서 어떤 키가 입력 되었는지 감지
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	// g.keys[:0] 키 목록을 비우고, 현재 눌려 있는 키들만으로 새로운 목록을 생성

	for _, key := range g.keys {
		switch key.String() {
		case "ArrowLeft":
			g_Character.AddPosX(-Speed) // left 키 입력시 속도만큼 x좌표 차감
		case "ArrowRight":
			g_Character.AddPosX(Speed) // right 키 입력시 속도만큼 x좌표 증가
		case "ArrowUp":
			g_Character.AddPosY(-Speed) // up 키 입력시 속도만큼 y좌표 차감
		case "ArrowDown":
			g_Character.AddPosY(Speed) // down 키 입력시 속도만큼 y좌표 증가
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

	// 0, 0 이미지의 출력 위치기준점으로 설정
	op.GeoM.Translate(float64(g_Character.GetPosX()), float64(g_Character.GetPosY()))

	screen.DrawImage(gopher_img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// 모든 게임 오브젝트가 포함 시켜야 할 위치 정보
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

// 좌표수정
func (o *ObjectPosInfo) AddPosX(x int) { o.x += x }
func (o *ObjectPosInfo) AddPosY(y int) { o.y += y }

// 캐릭터 객체
type Character struct {
	ObjectPosInfo
}

var (
	gopher_img *ebiten.Image
	wallImg    *ebiten.Image
	backImg    *ebiten.Image

	g_Character Character // <---- 캐릭터
	g_MapData   [][]int
)

const (
	screenWidth  = 500
	screenHeight = 500

	TypeWall = 1
	TypeBack = 2

	Speed                 = 2 //캐릭터 속도
	DefaultStartPositionX = 0 // <---- 추가
	DefaultStartPositionY = 0 // <----- 추가
)

func init() {
	// 캐릭터의 초기위치 설정
	g_Character.SetPosInfo(DefaultStartPositionX, DefaultStartPositionY)

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
