package party

import "fmt"

// CannotCreateAPartyWithoutName Error thrown when attempting to create a party without name
type CannotCreateAPartyWithoutName struct{}

func (err CannotCreateAPartyWithoutName) Error() string {
	return "Unable to create a party without name"
}

// CannotCreateAPartyWithAnAlreadyTakenName Error thrown when attempting to create a party with a name already taken
type CannotCreateAPartyWithAnAlreadyTakenName struct {
	name PartyName
}

func (err CannotCreateAPartyWithAnAlreadyTakenName) Error() string {
	return fmt.Sprintf("Unable to create the party, the name %v is already used", err.name)
}

// NoPartyNameProvided Error thrown when providing no party name
type NoPartyNameProvided struct{}

func (err NoPartyNameProvided) Error() string {
	return "No party name was provided"
}

// NotFound Error thrown when the party cannot be found
type NotFound struct {
	partyName PartyName
}

func (err NotFound) Error() string {
	return fmt.Sprintf("The party %v cannot be found", err.partyName)
}
