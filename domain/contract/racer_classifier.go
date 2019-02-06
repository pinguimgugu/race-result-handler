package contract

import "app/domain/entity"

type RacerClassifier interface {
	Make(map[string][]entity.RacePilotStatistic) []string
}
