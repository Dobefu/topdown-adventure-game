package fonts

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	FontDefaultMd *text.GoTextFace
	FontDefaultLg *text.GoTextFace
)

func init() {
	fontSrc, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

	if err != nil {
		log.Fatal(err)
	}

	FontDefaultMd = &text.GoTextFace{Source: fontSrc, Size: 24}
	FontDefaultLg = &text.GoTextFace{Source: fontSrc, Size: 32}
}
