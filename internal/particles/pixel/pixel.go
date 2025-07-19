package pixel

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/particles"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type Pixel struct {
	particles.Particle

	imgOptions *ebiten.DrawImageOptions
	img        *ebiten.Image
}

func NewPixel(position vectors.Vector3) (pixel *Pixel) {
	pixel = &Pixel{}
	pixel.SetPosition(position)

	pixel.imgOptions = &ebiten.DrawImageOptions{}
	pixel.img = ebiten.NewImage(1, 1)
	pixel.img.Fill(color.White)

	return pixel
}

func (p *Pixel) Update() (err error) {
	p.Particle.Update()

	return nil
}

func (p *Pixel) Draw(screen *ebiten.Image) {
	scene := *p.GetScene()
	camera := scene.GetCamera()
	position := p.GetPosition()

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(position.X, position.Y)

	camera.Draw(
		p.img,
		p.imgOptions,
		screen,
	)
}
