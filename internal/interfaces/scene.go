package interfaces

type Scene interface {
	Init()
	GetGameObjects() []GameObject
	AddGameObject(gameObject GameObject)
}
