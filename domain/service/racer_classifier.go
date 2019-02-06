package service

import (
	"app/domain/entity"
	"fmt"
	"sort"
	"time"
)

type RacerClassifier struct {
	MaxLapNumber    int
	punishmentByLap map[int]int64
}

func NewRacerClassifier() *RacerClassifier {
	return &RacerClassifier{
		4,
		map[int]int64{
			1: 999999999999999,
			2: 99999999999999,
			3: 9999999999999,
			4: 0,
		},
	}
}

func (rc *RacerClassifier) Make(resultByPilot map[string][]entity.RacePilotStatistic) []string {
	classification := map[int][]map[int64]string{}

	for pilotNumber, data := range resultByPilot {
		amountLapFinished := len(data)
		classification[amountLapFinished] = append(classification[amountLapFinished], map[int64]string{
			rc.getTotalRaceTimePilot(data): pilotNumber,
		})
	}

	return rc.orderClassification(classification)
}

func (rc *RacerClassifier) getTotalRaceTimePilot(data []entity.RacePilotStatistic) int64 {
	var totalTimeUnixFormat int64
	raceDay := time.Now()

	for _, lap := range data {
		totalTimeUnixFormat += rc.convertToUnixFormat(raceDay, lap.LapTime)
	}

	return totalTimeUnixFormat
}

func (rc *RacerClassifier) convertToUnixFormat(raceDay time.Time, lapTime string) int64 {
	layout := "20060102 15:04:05.000"
	stringDate := fmt.Sprintf("%d%02d%02d %02d:0%s",
		raceDay.Year(), raceDay.Month(), raceDay.Day(),
		raceDay.Hour(), lapTime)
	t2, _ := time.Parse(layout, stringDate)

	return t2.Unix()
}

func (rc *RacerClassifier) orderClassification(classification map[int][]map[int64]string) []string {
	var lapsOrdered []int
	var orderedLapsTime []int64
	var fineshedClassification []string
	timeByPilot := make(map[int64]string)
	for numberLaps, _ := range classification {
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
