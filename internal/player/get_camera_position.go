package player

import (
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	prevCameraPosition *vectors.Vector3
)

const (
	// MaxCameraOffset defines the maximum offset that the camera can have.
	MaxCameraOffset = 50
	// MaxCursorDistance defines the max distance that the cursor can be.
	MaxCursorDistance = 100
)

// GetCameraPosition gets the position that a camera should be following.
func (p *Player) GetCameraPosition() (position *vectors.Vector3) {
	if (*p.GetScene()).GetIsPaused() && p.state == state.StateDefault {
		return prevCameraPosition
	}

	playerPosition := *p.GetPosition()

	playerCenter := vectors.Vector3{
		X: playerPosition.X + (FRAME_WIDTH / 2),
		Y: playerPosition.Y + (FRAME_HEIGHT / 2),
		Z: 0,
	}

	if p.state != state.StateDefault {
		return &playerCenter
	}

	var cursorOffset vectors.Vector3

	if info, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok {
		cursorOffset = vectors.Vector3{
			X: info.Pos.X * MaxCameraOffset,
			Y: info.Pos.Y * MaxCameraOffset,
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

		screenDistanceSquared := distanceX*distanceX + distanceY*distanceY

		if screenDistanceSquared <= 0 {
			return &playerCenter
		}

		cameraWorldX := camera.X + centerX
		cameraWorldY := camera.Y + centerY
		cameraZoom := camera.ZoomFactor

		worldDistX := (cameraWorldX + distanceX/cameraZoom) - playerCenter.X
		worldDistY := (cameraWorldY + distanceY/cameraZoom) - playerCenter.Y
		worldDistanceSquared := worldDistX*worldDistX + worldDistY*worldDistY

		if worldDistanceSquared <= 0 {
			return &playerCenter
		}

		worldDistMagnitude := math.Sqrt(worldDistanceSquared)
		scale := math.Min(worldDistMagnitude/MaxCursorDistance, 1.0)

		cursorOffset = vectors.Vector3{
			X: (worldDistX / worldDistMagnitude) * scale * MaxCameraOffset,
			Y: (worldDistY / worldDistMagnitude) * scale * MaxCameraOffset,
			Z: 0,
		}
	}

	cameraPosition := playerCenter
	cameraPosition.Add(cursorOffset)

	prevCameraPosition = &cameraPosition

	return &cameraPosition
}
