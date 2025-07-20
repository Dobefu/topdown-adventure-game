package animation

type State int

const (
	StateOffsetIdle = iota * 8
	StateOffsetWalk
	StateOffsetRun
	StateOffsetAim
	StateOffsetShoot
	StateOffsetHurt
)

const (
	StateIdleRight = iota
	StateIdleDownRight
	StateIdleDown
	StateIdleDownLeft
	StateIdleLeft
	StateIdleUpLeft
	StateIdleUp
	StateIdleUpRight
	StateWalkRight
	StateWalkDownRight
	StateWalkDown
	StateWalkDownLeft
	StateWalkLeft
	StateWalkUpLeft
	StateWalkUp
	StateWalkUpRight
	StateRunRight
	StateRunDownRight
	StateRunDown
	StateRunDownLeft
	StateRunLeft
	StateRunUpLeft
	StateRunUp
	StateRunUpRight
	StateAimRight
	StateAimDownRight
	StateAimDown
	StateAimDownLeft
	StateAimLeft
	StateAimUpLeft
	StateAimUp
	StateAimUpRight
	StateShootRight
	StateShootDownRight
	StateShootDown
	StateShootDownLeft
	StateShootLeft
	StateShootUpLeft
	StateShootUp
	StateShootUpRight
	StateHurtRight
	StateHurtDownRight
	StateHurtDown
	StateHurtDownLeft
	StateHurtLeft
	StateHurtUpLeft
	StateHurtUp
	StateHurtUpRight
)
