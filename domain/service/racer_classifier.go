package service

import (
	"app/domain/entity"
	"app/infrastructure/formater"
	"sort"
)

type RacerClassifier struct {
	MaxLapNumber    int
	punishmentByLap map[int]int
}

func NewRacerClassifier() *RacerClassifier {
	return &RacerClassifier{
		4,
		map[int]int{
			1: 1000000000,
			2: 100000000,
			3: 10000000,
			4: 0,
		},
	}
}

func (rc *RacerClassifier) Make(resultByPilot map[string][]entity.RacePilotStatistic) []string {
	classification := map[int][]map[int]string{}

	for pilotNumber, data := range resultByPilot {
		amountLapFinished := len(data)
		classification[amountLapFinished] = append(classification[amountLapFinished], map[int]string{
			rc.getTotalRaceTimePilot(data): pilotNumber,
		})
	}

	return rc.orderClassification(classification)
}

func (rc *RacerClassifier) getTotalRaceTimePilot(data []entity.RacePilotStatistic) int {
	var totalTimeInt int

	for _, lap := range data {
		totalTimeInt += formater.ParseTimeToIntValue(lap.LapTime)
	}

	return totalTimeInt
}

func (rc *RacerClassifier) orderClassification(classification map[int][]map[int]string) []string {
	var lapsOrdered []int
	var orderedLapsTime []int
	var fineshedClassification []string
	timeByPilot := make(map[int]string)
	for numberLaps := range classification {
		lapsOrdered = append(lapsOrdered, numberLaps)
	}

	for _, lap := range lapsOrdered {
		for _, data := range classification[lap] {
			for timeRace, pilotNumber := range data {
				timeRaceWithPunish := timeRace + rc.punishmentByLap[lap]
				orderedLapsTime = append(orderedLapsTime, timeRaceWithPunish)
				timeByPilot[timeRaceWithPunish] = pilotNumber
			}
		}
	}
	sort.Slice(orderedLapsTime, func(i, j int) bool { return orderedLapsTime[i] < orderedLapsTime[j] })

	for _, raceTime := range orderedLapsTime {
		fineshedClassification = append(fineshedClassification, timeByPilot[raceTime])
	}

	return fineshedClassification
}
