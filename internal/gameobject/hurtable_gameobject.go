package gameobject

import (
	"log/slog"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

// HurtableGameObject defines a game object that can get damaged and can die.
type HurtableGameObject struct {
	CollidableGameObject

	health        int
	maxHealth     int
	deathCallback func()
}

// GetHealth gets the health of the game object.
func (h *HurtableGameObject) GetHealth() (health int) {
	return h.health
}

// SetHealth sets the health of the game object.
func (h *HurtableGameObject) SetHealth(health int) {
	h.health = health
}

// GetMaxHealth gets the max health of the game object.
func (h *HurtableGameObject) GetMaxHealth() (maxHealth int) {
	return h.maxHealth
}

// SetMaxHealth sets the max health of the game object.
func (h *HurtableGameObject) SetMaxHealth(maxHealth int) {
	h.maxHealth = maxHealth
}

// Damage handles damaging the game object.
func (h *HurtableGameObject) Damage(amount int, _ interfaces.GameObject) {
	h.health = int(math.Max(0, float64(h.health-amount)))

	if h.health <= 0 {
		if h.maxHealth == 0 {
			slog.Warn("maxHealth is zero. Did you forget to set it?")
		}

		if h.deathCallback != nil {
			h.deathCallback()
		}
	}
}

// Heal handles healing a game object.
func (h *HurtableGameObject) Heal(amount int) {
	h.health += amount
}

// GetDeathCallback gets the death callback function of the game object.
func (h *HurtableGameObject) GetDeathCallback() (callback func()) {
	return h.deathCallback
}

// SetDeathCallback sets the death callback function of the game object.
func (h *HurtableGameObject) SetDeathCallback(callback func()) {
	h.deathCallback = callback
}
