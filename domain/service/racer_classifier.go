package service

import (
	"app/domain/entity"
	"app/infrastructure/formater"
	"sort"
	"strconv"
)

type RacerClassifier struct {
	MaxLapNumber    int
	punishmentByLap map[string]int
}

func NewRacerClassifier() *RacerClassifier {
	return &RacerClassifier{
		4,
		map[string]int{
			"1": 1000000000,
			"2": 100000000,
			"3": 10000000,
			"4": 0,
		},
	}
}

func (rc *RacerClassifier) Make(resultByPilot map[string][]entity.RacePilotStatistic) []string {
	classification := []map[int]map[string]string{}

	for pilotNumber, data := range resultByPilot {
		classification = append(classification, map[int]map[string]string{
			rc.getTotalRaceTimePilot(data): map[string]string{
				"pilot_number": pilotNumber, "laps": strconv.Itoa(len(data)),
			},
		},
		)
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

func (rc *RacerClassifier) orderClassification(classification []map[int]map[string]string) []string {
	var orderedLapsTime []int
	var fineshedClassification []string
	timeByPilot := make(map[int]string)

	for _, data := range classification {
		for timeRace, pilotData := range data {
			timeRaceWithPunish := timeRace + rc.punishmentByLap[pilotData["laps"]]
			orderedLapsTime = append(orderedLapsTime, timeRaceWithPunish)
			timeByPilot[timeRaceWithPunish] = pilotData["pilot_number"]
		}
	}

	sort.Slice(orderedLapsTime, func(i, j int) bool { return orderedLapsTime[i] < orderedLapsTime[j] })

	for _, raceTime := range orderedLapsTime {
		fineshedClassification = append(fineshedClassification, timeByPilot[raceTime])
	}

	return fineshedClassification
}
