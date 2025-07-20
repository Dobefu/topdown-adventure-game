// Package fonts provides fonts in several sizes.
package fonts

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	//go:embed ttf/Oxanium-SemiBold.ttf
	fontSrc []byte

	// FontDefaultSm provides a small font.
	FontDefaultSm *text.GoTextFace
	// FontDefaultMd provides a medium size font.
	FontDefaultMd *text.GoTextFace
	// FontDefaultLg provides a large font.
	FontDefaultLg *text.GoTextFace
	// FontDefaultXl provides an extra large font.
	FontDefaultXl *text.GoTextFace
	// FontDefaultXxl provides an even larger font.
	FontDefaultXxl *text.GoTextFace
)

func init() {
	fontSrc, err := text.NewGoTextFaceSource(bytes.NewReader(fontSrc))

	if err != nil {
		log.Fatal(err)
	}

	FontDefaultSm = &text.GoTextFace{Source: fontSrc, Size: 12}
	FontDefaultMd = &text.GoTextFace{Source: fontSrc, Size: 16}
	FontDefaultLg = &text.GoTextFace{Source: fontSrc, Size: 24}
	FontDefaultXl = &text.GoTextFace{Source: fontSrc, Size: 32}
	FontDefaultXxl = &text.GoTextFace{Source: fontSrc, Size: 48}
}
