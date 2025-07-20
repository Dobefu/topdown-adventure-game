// Package state provides constants for generic states.
package state

// State houses the different state constants.
type State byte

const (
	// StateDefault is the default state.
	StateDefault = iota
	// StateHurt is the hurt state.
	StateHurt
)
