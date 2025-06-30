package player

import (
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

const MAX_CAMERA_OFFSET = 50
const MAX_CURSOR_DISTANCE = 100

func (p *Player) GetCameraPosition() (position *vectors.Vector3) {
	playerPosition := *p.GetPosition()

	playerCenter := vectors.Vector3{
		X: playerPosition.X + (FRAME_WIDTH / 2),
		Y: playerPosition.Y + (FRAME_HEIGHT / 2),
		Z: 0,
	}

	var cursorOffset vectors.Vector3

	if info, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok {
		cursorOffset = vectors.Vector3{
			X: info.Pos.X * MAX_CAMERA_OFFSET,
			Y: info.Pos.Y * MAX_CAMERA_OFFSET,
			Z: 0,
		}
	} else if p.input.ActionIsPressed(input.ActionAimMouse) {
		cursorX, cursorY := ebiten.CursorPosition()
		scene := *p.GetScene()
		camera := scene.GetCamera()

		centerX := camera.Width / 2
		centerY := camera.Height / 2
		distanceX := float64(cursorX) - centerX
		distanceY := float64(cursorY) - centerY

		screenDistance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)

		if screenDistance > 0 {
			scale := math.Min(screenDistance/MAX_CURSOR_DISTANCE, 1.0)
			cursorOffset = vectors.Vector3{
				X: (distanceX / screenDistance) * scale * MAX_CAMERA_OFFSET,
				Y: (distanceY / screenDistance) * scale * MAX_CAMERA_OFFSET,
				Z: 0,
			}
		}
	}

	cameraPosition := playerCenter
	cameraPosition.Add(cursorOffset)

	return &cameraPosition
}
