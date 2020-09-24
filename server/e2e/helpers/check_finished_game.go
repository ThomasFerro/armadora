package helpers

func CheckFinishedGame(partyId string) error {
	_, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	// TODO: Check finished game attributes
	return nil
}
