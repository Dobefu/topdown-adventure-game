package fonts

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	FontDefaultSm  *text.GoTextFace
	FontDefaultMd  *text.GoTextFace
	FontDefaultLg  *text.GoTextFace
	FontDefaultXl  *text.GoTextFace
	FontDefaultXxl *text.GoTextFace
)

func init() {
	fontSrc, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

	if err != nil {
		log.Fatal(err)
	}

	FontDefaultSm = &text.GoTextFace{Source: fontSrc, Size: 16}
	FontDefaultMd = &text.GoTextFace{Source: fontSrc, Size: 24}
	FontDefaultLg = &text.GoTextFace{Source: fontSrc, Size: 32}
	FontDefaultXl = &text.GoTextFace{Source: fontSrc, Size: 48}
	FontDefaultXxl = &text.GoTextFace{Source: fontSrc, Size: 64}
}
