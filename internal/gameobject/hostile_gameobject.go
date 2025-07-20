package gameobject

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

// HostileGameObject defines a game object that can hurt a player.
type HostileGameObject struct {
	interfaces.HostileGameObject
}
