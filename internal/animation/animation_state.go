package animation

type AnimationState int

const (
	AnimationStateIdleUp = iota
	AnimationStateIdleDown
	AnimationStateIdleLeft
	AnimationStateIdleRight
	AnimationStateWalkingUp
	AnimationStateWalkingDown
	AnimationStateWalkingLeft
	AnimationStateWalkingRight
)
