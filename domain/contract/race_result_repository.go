package contract

import "app/domain/entity"

type RaceResultRepository interface {
	GetList() []entity.RacePilotStatistic
	CreateClassification([]string, map[string]entity.RacePilotStatistic)
}
