package game_object

import (
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

type HurtableGameObject struct {
	interfaces.HurtableGameObject
	CollidableGameObject

	health        int
	maxHealth     int
	deathCallback func()
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
	h.health = int(math.Max(0, float64(h.health-amount)))

	if h.health <= 0 {
		if h.maxHealth == 0 {
			log.Println("maxHealth is zero. Did you forget to set it?")
		}

		if h.deathCallback != nil {
			h.deathCallback()
		}
	}
}

func (h *HurtableGameObject) Heal(amount int) {
	h.health += amount
}

func (h *HurtableGameObject) GetDeathCallback() (callback func()) {
	return h.deathCallback
}

func (h *HurtableGameObject) SetDeathCallback(callback func()) {
	h.deathCallback = callback
}
