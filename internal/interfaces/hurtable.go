package interfaces

type Hurtable interface {
	GetHealth() (health int)
	SetHealth(health int)
	Damage(amount int)
	Heal(amount int)
}
