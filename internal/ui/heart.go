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

	//go:embed img/heart-empty.png
	heartEmptyImgBytes []byte
	heartEmptyImg      *ebiten.Image
	heartEmptyImgLeft  *ebiten.Image
	heartEmptyImgRight *ebiten.Image

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

	img, _, err = image.Decode(bytes.NewReader(heartEmptyImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	heartEmptyImg = ebiten.NewImageFromImage(img)
	heartEmptyImgLeft = ebiten.NewImageFromImage(
		heartEmptyImg.SubImage(image.Rect(0, 0, heartImgWidth/2, heartImgHeight)),
	)
	heartEmptyImgRight = ebiten.NewImageFromImage(
		heartEmptyImg.SubImage(image.Rect(8, 0, heartImgWidth, heartImgHeight)),
	)
}

func DrawHealthBar(
	screen *ebiten.Image,
	position vectors.Vector2,
	health int,
	maxHealth int,
) {
	op := &ebiten.DrawImageOptions{}
	i := 0

	for range health {
		op.GeoM.Reset()
		op.GeoM.Translate(position.X+float64(i*heartImgWidth/2), position.Y)

		if i%2 == 0 {
			screen.DrawImage(heartFilledImgLeft, op)
		} else {
			screen.DrawImage(heartFilledImgRight, op)
		}

		i += 1
	}

	for range maxHealth - health {
		op.GeoM.Reset()
		op.GeoM.Translate(position.X+float64(i*heartImgWidth/2), position.Y)

		if i%2 == 0 {
			screen.DrawImage(heartEmptyImgLeft, op)
		} else {
			screen.DrawImage(heartEmptyImgRight, op)
		}

		i += 1
	}
}
