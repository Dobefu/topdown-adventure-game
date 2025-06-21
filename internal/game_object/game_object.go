package game_object

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
)

type GameObject struct {
	interfaces.GameObject

	position vectors.Vector3
	isActive bool
}
