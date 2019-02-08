package service

import (
	"app/domain/entity"
	"app/infrastructure/formater"
	"sort"
)

type IdentifierBestLap struct {
	resultRace []entity.RacePilotStatistic
}

func (bl *IdentifierBestLap) Attach(pilotResult entity.RacePilotStatistic) {
	bl.resultRace = append(bl.resultRace, pilotResult)
}

func (bl *IdentifierBestLap) Get() (string, string) {
	if len(bl.resultRace) == 0 {
		return "", ""
	}

	pilot := make(map[int]map[string]string, 0)
	sortLap := []int{}

	for _, data := range bl.resultRace {
		timeInt := formater.ParseTimeToIntValue(data.BestLap)
		pilot[timeInt] = map[string]string{
			"pilot_number": data.Number,
			"lap_time":     data.BestLap,
		}
		sortLap = append(sortLap, timeInt)
	}

	sort.Ints(sortLap)

	bestLap := sortLap[0]

	return pilot[bestLap]["pilot_number"], pilot[bestLap]["lap_time"]
}
