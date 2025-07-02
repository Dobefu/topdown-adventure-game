package interfaces

type HurtableGameObject interface {
	GetHealth() (health int)
	SetHealth(health int)
	GetMaxHealth() (maxHealth int)
	Damage(amount int)
	Heal(amount int)
}
