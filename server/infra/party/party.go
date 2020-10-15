package party

// Restriction The restriction for joining a party
type Restriction string

const (
	// Public A public party, anyone can see it
	Public Restriction = "PUBLIC"
	// Private A private party, only the one knowing the identifier can see it
	Private Restriction = "PRIVATE"
)

// PartyID A party identifier
type PartyID string

// Party A party
type Party struct {
	Restriction Restriction
	Name        string
}

// NewParty Create a new party
func NewParty(partyName string, restriction Restriction) Party {
	return Party{
		Name:        partyName,
		Restriction: restriction,
	}
}
