package dto

import (
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/score"
	"github.com/ThomasFerro/armadora/game/warrior"
)

type EventDto interface{}

type GameCreatedDto struct{}

type PlayerJoinedDto struct {
	Nickname  string `json:"nickname" bson:"nickname"`
	Character int    `json:"character" bson:"character"`
}

type GameStartedDto struct{}

type GoldStacksDistributedDto struct {
	GoldStacks []int `json:"gold_stacks" bson:"gold_stacks"`
}

type WarriorsDistributedDto struct {
	WarriorsDistributed WarriorsDto `json:"warriors" bson:"warriors"`
}

type PalisadesDistributedDto struct {
	Count int `json:"count" bson:"count"`
}

type NextPlayerDto struct{}

type PalisadePutDto struct {
	Player int `json:"player" bson:"player"`
	X1     int `json:"x1" bson:"x1"`
	Y1     int `json:"y1" bson:"y1"`
	X2     int `json:"x2" bson:"x2"`
	Y2     int `json:"y2" bson:"y2"`
}

type WarriorPutDto struct {
	Player   int `json:"player" bson:"player"`
	Strength int `json:"strength" bson:"strength"`
	X        int `json:"x" bson:"x"`
	Y        int `json:"y" bson:"y"`
}

type TurnPassedDto struct {
	Player int `json:"player" bson:"player"`
}

type GameFinishedDto struct {
	Scores ScoresDto `json:"scores" bson:"scores"`
}

func ToEventsDto(events []event.Event) []EventDto {
	mappedEvents := []EventDto{}
	for _, nextEvent := range events {
		if mappedNextEvent := ToEventDto(nextEvent); mappedNextEvent != nil {
			mappedEvents = append(mappedEvents, mappedNextEvent)
		}
	}
	return mappedEvents
}

func ToEventDto(nextEvent event.Event) EventDto {
	switch typedEvent := nextEvent.(type) {
	case event.GameCreated:
		return GameCreatedDto{}
	case event.PlayerJoined:
		return PlayerJoinedDto{
			Nickname:  typedEvent.Nickname,
			Character: int(typedEvent.Character),
		}
	case event.GameStarted:
		return GameStartedDto{}
	case event.GoldStacksDistributed:
		return GoldStacksDistributedDto{
			GoldStacks: typedEvent.GoldStacks,
		}
	case event.WarriorsDistributed:
		return WarriorsDistributedDto{
			WarriorsDistributed: toWarriorsDto(typedEvent.WarriorsDistributed),
		}
	case event.PalisadesDistributed:
		return PalisadesDistributedDto{
			Count: typedEvent.Count,
		}
	case event.NextPlayer:
		return NextPlayerDto{}
	case event.PalisadePut:
		return PalisadePutDto{
			Player: typedEvent.Player,
			X1:     typedEvent.X1,
			Y1:     typedEvent.Y1,
			X2:     typedEvent.X2,
			Y2:     typedEvent.Y2,
		}
	case event.WarriorPut:
		return WarriorPutDto{
			Player:   typedEvent.Player,
			Strength: typedEvent.Strength,
			X:        typedEvent.Position.X,
			Y:        typedEvent.Position.Y,
		}
	case event.TurnPassed:
		return TurnPassedDto{
			Player: typedEvent.Player,
		}
	case event.GameFinished:
		return GameFinishedDto{
			Scores: toScoresDto(typedEvent.Scores),
		}
	}
	return nil
}

func FromEventsDto(eventsDto []EventDto) []event.Event {
	mappedEvents := []event.Event{}
	for _, nextEvent := range eventsDto {
		if mappedNextEvent := FromEventDto(nextEvent); mappedNextEvent != nil {
			mappedEvents = append(mappedEvents, mappedNextEvent)
		}
	}
	return mappedEvents
}

func FromEventDto(eventDto EventDto) event.Event {
	switch typedEvent := eventDto.(type) {
	case GameCreatedDto:
		return event.GameCreated{}
	case PlayerJoinedDto:
		return event.PlayerJoined{
			Nickname:  typedEvent.Nickname,
			Character: character.Character(typedEvent.Character),
		}
	case GameStartedDto:
		return event.GameStarted{}
	case GoldStacksDistributedDto:
		return event.GoldStacksDistributed{
			GoldStacks: typedEvent.GoldStacks,
		}
	case WarriorsDistributedDto:
		return event.WarriorsDistributed{
			WarriorsDistributed: fromWarriorsDto(typedEvent.WarriorsDistributed),
		}
	case PalisadesDistributedDto:
		return event.PalisadesDistributed{
			Count: typedEvent.Count,
		}
	case NextPlayerDto:
		return event.NextPlayer{}
	case PalisadePutDto:
		return event.PalisadePut{
			Player: typedEvent.Player,
			X1:     typedEvent.X1,
			Y1:     typedEvent.Y1,
			X2:     typedEvent.X2,
			Y2:     typedEvent.Y2,
		}
	case WarriorPutDto:
		return event.WarriorPut{
			Player:   typedEvent.Player,
			Strength: typedEvent.Strength,
			Position: board.Position{
				X: typedEvent.X,
				Y: typedEvent.Y,
			},
		}
	case TurnPassedDto:
		return event.TurnPassed{
			Player: typedEvent.Player,
		}
	case GameFinishedDto:
		return event.GameFinished{
			Scores: fromScoresDto(typedEvent.Scores),
		}
	}
	return nil
}

func fromWarriorsDto(warriorsDto WarriorsDto) warrior.Warriors {
	return warrior.NewWarriors(
		warriorsDto.OnePoint,
		warriorsDto.TwoPoints,
		warriorsDto.ThreePoints,
		warriorsDto.FourPoints,
		warriorsDto.FivePoints,
	)
}

func fromScoreDto(scoreDto ScoreDto) score.Score {
	return score.NewScore(
		scoreDto.Player,
		scoreDto.GoldStacks,
	)
}

func fromScoresDto(scoresDto ScoresDto) score.Scores {
	mappedScores := score.Scores{}

	if scoresDto != nil {
		for rank, score := range scoresDto {
			mappedScores[rank] = fromScoreDto(score)
		}
	}

	return mappedScores
}
