package interfaces

type HurtableGameObject interface {
	GetHealth() (health int)
	SetHealth(health int)
	GetMaxHealth() (maxHealth int)
	SetMaxHealth(maxHealth int)
	Damage(amount int, source GameObject)
	Heal(amount int)
	GetDeathCallback() (callback func())
	SetDeathCallback(callback func())
}
