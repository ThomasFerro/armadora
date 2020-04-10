package event

// Event A game event
type Event interface {
	EventMessage() string
}
