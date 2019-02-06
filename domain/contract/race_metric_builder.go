package contract

import "app/domain/entity"

type RaceMetricBuilder interface {
	Build(resultList []entity.RacePilotStatistic) entity.RacePilotStatistic
}
