// Package gameobject provides various types of game objects.
package gameobject

import (
	"sync/atomic"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	nextGameObjectID uint64
)

// GameObject is the base game object type.
// It can be used directly, and other game object types may be embedded
// to expand its functionality.
type GameObject struct {
	ID uint64

	scene    *interfaces.Scene
	Position vectors.Vector3
	isActive bool
}

// Init initializes the game object.
func (g *GameObject) Init() {
	if g.ID == 0 {
		g.ID = atomic.AddUint64(&nextGameObjectID, 1)
	}
}

// GetID gets the ID of the game object.
func (g *GameObject) GetID() (id uint64) {
	return g.ID
}

// Draw draws the game object.
// This should be overridden on structs that embed this one.
func (g *GameObject) Draw(_ *ebiten.Image) {
	// noop
}

// DrawBelow draws below the game object.
// This should be overridden on structs that embed this one.
func (g *GameObject) DrawBelow(_ *ebiten.Image) {
	// noop
}

// DrawAbove draws above the game object.
// This should be overridden on structs that embed this one.
func (g *GameObject) DrawAbove(_ *ebiten.Image) {
	// noop
}

// DrawUI draws on the UI layer.
// This should be overridden on structs that embed this one.
func (g *GameObject) DrawUI(_ *ebiten.Image) {
	// noop
}

// Update runs during the update function of the game.
func (g *GameObject) Update() (err error) {
	return nil
}

// Damage handles damaging a game object.
func (g *GameObject) Damage(amount int, source interfaces.GameObject) {
	// noop
}
