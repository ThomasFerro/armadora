package party

// Restriction The restriction for joining a party
type Restriction string

const (
	// Public A public party, anyone can see it
	Public Restriction = "PUBLIC"
	// Private A private party, only the one knowing the identifier can see it
	Private Restriction = "PRIVATE"
)

// Status The party's status
type Status string

const (
	// Open An open party
	Open Status = "OPEN"
	// Close A closed party
	Close Status = "CLOSE"
)

// PartyName The party name
type PartyName string

// Party A party
type Party struct {
	Name        PartyName
	Restriction Restriction
	Status      Status
}

// Close Return the closed party
func (p Party) Close() Party {
	p.Status = Close
	return p
}

// NewParty Create a new party
func NewParty(partyName PartyName, restriction Restriction, status Status) Party {
	return Party{
		Name:        partyName,
		Restriction: restriction,
		Status:      Open,
	}
}
