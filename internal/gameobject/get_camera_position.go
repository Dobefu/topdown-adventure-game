package gameobject

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
	// MaxCameraCursorDistance defines the max distance that the cursor can be.
	MaxCameraCursorDistance = 100
)

// GetCameraPosition gets the position that a camera should be following.
func (g *GameObject) GetCameraPosition() (position *vectors.Vector3) {
	if (*g.GetScene()).GetIsPaused() && g.State == state.StateDefault {
		return prevCameraPosition
	}

	playerCenter := vectors.Vector3{
		X: g.Position.X + (float64(g.FrameWidth) / 2),
		Y: g.Position.Y + (float64(g.FrameHeight) / 2),
		Z: 0,
	}

	if g.State != state.StateDefault {
		return &playerCenter
	}

	var cursorOffset vectors.Vector3

	if g.Input != nil {
		if info, ok := g.Input.PressedActionInfo(input.ActionAimAnalog); ok {
			cursorOffset = vectors.Vector3{
				X: info.Pos.X * MaxCameraOffset,
				Y: info.Pos.Y * MaxCameraOffset,
				Z: 0,
			}
		} else if g.Input.ActionIsPressed(input.ActionAimMouse) {
			cursorX, cursorY := ebiten.CursorPosition()
			scene := *g.GetScene()
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
			scale := math.Min(worldDistMagnitude/MaxCameraCursorDistance, 1.0)

			cursorOffset = vectors.Vector3{
				X: (worldDistX / worldDistMagnitude) * scale * MaxCameraOffset,
				Y: (worldDistY / worldDistMagnitude) * scale * MaxCameraOffset,
				Z: 0,
			}
		}
	}

	cameraPosition := playerCenter
	cameraPosition.Add(cursorOffset)

	prevCameraPosition = &cameraPosition

	return &cameraPosition
}
