package infra

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/warrior"
)

type StateDto string

type CharacterDto string

type WarriorsDto struct {
	OnePoint    int `json:"one_point"`
	TwoPoints   int `json:"two_points"`
	ThreePoints int `json:"three_points"`
	FourPoints  int `json:"four_points"`
	FivePoints  int `json:"five_points"`
}

type PlayerDto struct {
	Nickname  string       `json:"nickname"`
	Character CharacterDto `json:"character"`
	Warriors  WarriorsDto  `json:"warriors"`
}

type BoardDto struct {
}

type GameDto struct {
	State               StateDto    `json:"state"`
	Players             []PlayerDto `json:"players"`
	CurrentPlayer       int         `json:"current_player"`
	Board               BoardDto    `json:"board"`
	AvailableCharacters []string    `json:"available_characters"`
}

func toStateDto(state game.State) StateDto {
	switch state {
	case game.WaitingForPlayers:
		return "WaitingForPlayers"
	case game.Started:
		return "Started"
	}
	return ""
}

func toCharacterDto(characterToMap character.Character) CharacterDto {
	switch characterToMap {
	case character.Orc:
		return "Orc"
	case character.Goblin:
		return "Goblin"
	case character.Elf:
		return "Elf"
	case character.Mage:
		return "Mage"
	}
	return ""
}

func toWarriorsDto(warriors warrior.Warriors) WarriorsDto {
	if warriors == nil {
		return WarriorsDto{}
	}
	return WarriorsDto{
		OnePoint:    warriors.OnePoint(),
		TwoPoints:   warriors.TwoPoints(),
		ThreePoints: warriors.ThreePoints(),
		FourPoints:  warriors.FourPoints(),
		FivePoints:  warriors.FivePoints(),
	}
}

func toPlayersDto(players []game.Player) []PlayerDto {
	playersDto := []PlayerDto{}
	for _, player := range players {
		playersDto = append(playersDto, PlayerDto{
			Nickname:  player.Nickname(),
			Character: toCharacterDto(player.Character()),
			Warriors:  toWarriorsDto(player.Warriors()),
		})
	}
	return playersDto
}

func getAvailableCharacters(players []PlayerDto) []string {
	availableCharactersMap := map[CharacterDto]bool{
		"Orc":    true,
		"Goblin": true,
		"Elf":    true,
		"Mage":   true,
	}
	for _, player := range players {
		availableCharactersMap[player.Character] = false
	}
	availableCharacters := []string{}
	for character, available := range availableCharactersMap {
		if available {
			availableCharacters = append(availableCharacters, string(character))
		}
	}
	return availableCharacters
}

func ToGameDto(game game.Game) GameDto {
	// TODO: Map the board
	playersDto := toPlayersDto(game.Players())
	return GameDto{
		State:               toStateDto(game.State()),
		Players:             playersDto,
		CurrentPlayer:       game.CurrentPlayer(),
		AvailableCharacters: getAvailableCharacters(playersDto),
	}
}
