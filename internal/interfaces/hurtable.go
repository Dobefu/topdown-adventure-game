package interfaces

// Hurtable defines the interface for objects that can be damaged.
type Hurtable interface {
	Damage(amount int, source GameObject)
	GetHealth() int
	SetHealth(health int)
	GetMaxHealth() int
	SetMaxHealth(maxHealth int)
}
