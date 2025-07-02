package game_object

import "github.com/Dobefu/topdown-adventure-game/internal/interfaces"

type HurtableGameObject struct {
	interfaces.HurtableGameObject

	health int
}

func (h *HurtableGameObject) GetHealth() (health int) {
	return health
}

func (h *HurtableGameObject) SetHealth(health int) {
	h.health = health
}

func (h *HurtableGameObject) Damage(amount int) {
	h.health -= amount
}

func (h *HurtableGameObject) Heal(amount int) {
	h.health += amount
}
