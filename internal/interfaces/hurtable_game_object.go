package interfaces

type HurtableGameObject interface {
	GetHealth() (health int)
	SetHealth(health int)
	Damage(amount int)
	Heal(amount int)
}
