package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 400
	screenHeight = 450
)

type Slider struct {
	x, y, w, h float32
	min, max   float32
	value      float32
	knobX      float32
}

func NewSlider(x, y, w, h, min, max, initial float32) *Slider {
	s := &Slider{x: x, y: y, w: w, h: h, min: min, max: max, value: initial}
	s.setValue(initial)
	return s
}

func (s *Slider) setValue(val float32) {
	s.value = val
	ratio := (val - s.min) / (s.max - s.min)
	padding := float32(10)
	sliderW := s.w - (padding * 2)
	s.knobX = s.x + padding + (ratio * sliderW)
}

func (s *Slider) Update(mx, my int, isPressed bool) {
	if isPressed && float32(mx) >= s.x && float32(mx) <= s.x+s.w && float32(my) >= s.y && float32(my) <= s.y+s.h {
		padding := float32(10)
		left := s.x + padding
		right := s.x + s.w - padding
		
		x := float32(mx)
		if x < left { x = left }
		if x > right { x = right }
		
		s.knobX = x
		ratio := (x - left) / (s.w - (padding * 2))
		s.value = s.min + ratio*(s.max-s.min)
	}
}


// 57行目付近の修正
func (s *Slider) Draw(screen *ebiten.Image) {
	// 背景
	vector.DrawFilledRect(screen, s.x, s.y, s.w, s.h, color.RGBA{225, 225, 225, 255}, true)
	// バー
	vector.DrawFilledRect(screen, s.x+10, s.y+10, s.w-20, s.h-20, color.RGBA{64, 64, 128, 255}, true)
	// ツマミ
	vector.DrawFilledRect(screen, s.knobX-2, s.y, 4, s.h, color.RGBA{0, 0, 255, 255}, true)
}

type Game struct {
	slider    *Slider
	isPressed bool
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.isPressed = true
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.isPressed = false
	}
	mx, my := ebiten.CursorPosition()
	g.slider.Update(mx, my, g.isPressed)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.slider.Draw(screen)

	// グリッド線の描画
	for i := 0; i < 400; i += 10 {
		c := color.RGBA{225, 225, 225, 255}
		vector.StrokeLine(screen, 0, float32(i), 400, float32(i), 1, c, true)
		vector.StrokeLine(screen, float32(i), 0, float32(i), 400, 1, c, true)
	}
	// 軸
	vector.StrokeLine(screen, 0, 200, 400, 200, 1.5, color.Black, true)
	vector.StrokeLine(screen, 200, 0, 200, 400, 1.5, color.Black, true)
	// 単位円
	vector.StrokeCircle(screen, 200, 200, 150, 2, color.RGBA{255, 0, 0, 255}, true)

	// 三角関数の計算
	theta := float64(g.slider.value)
	rad := theta * math.Pi / 180.0
	cosV := math.Cos(rad)
	sinV := math.Sin(rad)

	xpos := float32(cosV*150 + 200)
	ypos := float32(-sinV*150 + 200) // Y軸反転

	// ベクトル線の描画
	vector.StrokeLine(screen, xpos, ypos, xpos, 200, 2, color.RGBA{0, 0, 192, 255}, true)  // cos (X軸への垂線)
	vector.StrokeLine(screen, xpos, ypos, 200, ypos, 2, color.RGBA{0, 192, 0, 255}, true)  // sin (Y軸への垂線)
	vector.StrokeLine(screen, xpos, ypos, 200, 200, 2, color.RGBA{192, 0, 0, 255}, true)  // 半径

	// テキスト出力
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("cos:%.3f", cosV), int(xpos), 200)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("sin:%.3f", sinV), 200, int(ypos))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Trigonometric Function Simulator")
	g := &Game{slider: NewSlider(20, 410, 360, 35, 0, 360, 0)}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
