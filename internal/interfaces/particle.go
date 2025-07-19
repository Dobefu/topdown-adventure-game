package interfaces

type Particle interface {
	Update()
	GetLifetime() (lifetime int)
	SetLifetime(lifetime int)
}
