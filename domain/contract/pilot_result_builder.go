package contract

import "app/domain/entity"

type PilotResultBuilder interface {
	Build(resultList []entity.RacePilotStatistic) entity.RacePilotStatistic
}
