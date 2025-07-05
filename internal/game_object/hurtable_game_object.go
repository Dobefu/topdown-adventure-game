package game_object

import (
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

type HurtableGameObject struct {
	interfaces.HurtableGameObject
	GameObject

	health    int
	maxHealth int
}

func (h *HurtableGameObject) GetHealth() (health int) {
	return h.health
}

func (h *HurtableGameObject) SetHealth(health int) {
	h.health = health
}

func (h *HurtableGameObject) GetMaxHealth() (maxHealth int) {
	return h.maxHealth
}

func (h *HurtableGameObject) SetMaxHealth(maxHealth int) {
	h.maxHealth = maxHealth
}

func (h *HurtableGameObject) Damage(amount int) {
	h.health -= amount

	if h.health <= 0 {
		if h.maxHealth == 0 {
			log.Println("maxHealth is zero. Did you forget to set it?")
		}

		h.Die()
	}
}

func (h *HurtableGameObject) Heal(amount int) {
	h.health += amount
}

func (h *HurtableGameObject) Die() {
	h.SetIsActive(false)
}
