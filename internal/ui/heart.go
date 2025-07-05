package ui

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed img/heart-filled.png
	heartFilledImgBytes []byte
	heartFilledImg      *ebiten.Image
	heartFilledImgLeft  *ebiten.Image
	heartFilledImgRight *ebiten.Image

	heartImgWidth  int
	heartImgHeight int
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(heartFilledImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	heartImgWidth = img.Bounds().Max.X
	heartImgHeight = img.Bounds().Max.X

	heartFilledImg = ebiten.NewImageFromImage(img)
	heartFilledImgLeft = ebiten.NewImageFromImage(
		heartFilledImg.SubImage(image.Rect(0, 0, heartImgWidth/2, heartImgHeight)),
	)
	heartFilledImgRight = ebiten.NewImageFromImage(
		heartFilledImg.SubImage(image.Rect(8, 0, heartImgWidth, heartImgHeight)),
	)
}

func DrawHealthBar(screen *ebiten.Image, position vectors.Vector2, health int) {
	op := &ebiten.DrawImageOptions{}

	for i := range health {
		op.GeoM.Reset()
		op.GeoM.Translate(position.X+float64(i*heartImgWidth/2), position.Y)

		if i%2 == 0 {
			screen.DrawImage(heartFilledImgLeft, op)
		} else {
			screen.DrawImage(heartFilledImgRight, op)
		}
	}
}
