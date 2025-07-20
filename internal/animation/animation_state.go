package animation

// State handles an animation state.
type State byte

const (
	// StateOffsetIdle animation state offset.
	StateOffsetIdle = iota * 8
	// StateOffsetWalk animation state offset.
	StateOffsetWalk
	// StateOffsetRun animation state offset.
	StateOffsetRun
	// StateOffsetAim animation state offset.
	StateOffsetAim
	// StateOffsetShoot animation state offset.
	StateOffsetShoot
	// StateOffsetHurt animation state offset.
	StateOffsetHurt
)

const (
	// StateIdleRight animation state index.
	StateIdleRight = iota
	// StateIdleDownRight animation state index.
	StateIdleDownRight
	// StateIdleDown animation state index.
	StateIdleDown
	// StateIdleDownLeft animation state index.
	StateIdleDownLeft
	// StateIdleLeft animation state index.
	StateIdleLeft
	// StateIdleUpLeft animation state index.
	StateIdleUpLeft
	// StateIdleUp animation state index.
	StateIdleUp
	// StateIdleUpRight animation state index.
	StateIdleUpRight
	// StateWalkRight animation state index.
	StateWalkRight
	// StateWalkDownRight animation state index.
	StateWalkDownRight
	// StateWalkDown animation state index.
	StateWalkDown
	// StateWalkDownLeft animation state index.
	StateWalkDownLeft
	// StateWalkLeft animation state index.
	StateWalkLeft
	// StateWalkUpLeft animation state index.
	StateWalkUpLeft
	// StateWalkUp animation state index.
	StateWalkUp
	// StateWalkUpRight animation state index.
	StateWalkUpRight
	// StateRunRight animation state index.
	StateRunRight
	// StateRunDownRight animation state index.
	StateRunDownRight
	// StateRunDown animation state index.
	StateRunDown
	// StateRunDownLeft animation state index.
	StateRunDownLeft
	// StateRunLeft animation state index.
	StateRunLeft
	// StateRunUpLeft animation state index.
	StateRunUpLeft
	// StateRunUpRight animation state index.
	StateRunUpRight
	// StateAimRight animation state index.
	StateAimRight
	// StateAimDownRight animation state index.
	StateAimDownRight
	// StateAimDown animation state index.
	StateAimDown
	// StateAimDownLeft animation state index.
	StateAimDownLeft
	// StateAimLeft animation state index.
	StateAimLeft
	// StateAimUpLeft animation state index.
	StateAimUpLeft
	// StateAimUp animation state index.
	StateAimUp
	// StateAimUpRight animation state index.
	StateAimUpRight
	// StateShootRight animation state index.
	StateShootRight
	// StateShootDownRight animation state index.
	StateShootDownRight
	// StateShootDown animation state index.
	StateShootDown
	// StateShootDownLeft animation state index.
	StateShootDownLeft
	// StateShootLeft animation state index.
	StateShootLeft
	// StateShootUpLeft animation state index.
	StateShootUpLeft
	// StateShootUp animation state index.
	StateShootUp
	// StateShootUpRight animation state index.
	StateShootUpRight
	// StateHurtRight animation state index.
	StateHurtRight
	// StateHurtDownRight animation state index.
	StateHurtDownRight
	// StateHurtDown animation state index.
	StateHurtDown
	// StateHurtDownLeft animation state index.
	StateHurtDownLeft
	// StateHurtLeft animation state index.
	StateHurtLeft
	// StateHurtUpLeft animation state index.
	StateHurtUpLeft
	// StateHurtUp animation state index.
	StateHurtUp
	// StateHurtUpRight animation state index.
	StateHurtUpRight
)
