package interfaces

type Scene interface {
	Init()
	InitUI()
	GetGameObjects() []GameObject
	AddGameObject(gameObject GameObject)
}
