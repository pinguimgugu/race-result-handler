package service

import (
	"app/domain/entity"
	"app/infrastructure/formater"
	"sort"

	"math"
	"regexp"
	"strconv"
	"strings"
)

type PilotResultBuilder struct {
}

func NewPilotResultBuilder() *PilotResultBuilder {
	return &PilotResultBuilder{}
}

func (rm *PilotResultBuilder) Build(resultList []entity.RacePilotStatistic) entity.RacePilotStatistic {
	var miliseconds int
	var totalSecondsDuration int
	var speedAverage float64
	var speedAmount float64
	allLapTime := make(map[int]string)

	pilotStatistic := entity.RacePilotStatistic{
		LapAmount: len(resultList),
	}

	for _, data := range resultList {
		pilotStatistic.Name = data.Name
		pilotStatistic.Number = data.Number
		minutes, seconds, milisecondsInt := rm.parseLapTime(data.LapTime)
		f, _ := strconv.ParseFloat(
			strings.Replace(data.SpeedLapAverage, ",", ".", -1), 64,
		)
		totalSecondsDuration += (seconds + 60*minutes)
		speedAmount += f
		miliseconds += milisecondsInt
		allLapTime[formater.ParseTimeToIntValue(data.LapTime)] = data.LapTime
	}

	speedAverage = speedAmount / float64(len(resultList))
	totalSecondsDuration += int(math.Round(float64(miliseconds) / float64(1000)))

	pilotStatistic.SpeedRaceAverage = speedAverage
	pilotStatistic.RaceTime = formater.FormatBySeconds(totalSecondsDuration)
	pilotStatistic.BestLap = rm.getBestLap(allLapTime)

	return pilotStatistic
}

func (rm *PilotResultBuilder) parseLapTime(lapTime string) (int, int, int) {
	rp := regexp.MustCompile("[0-9]+")
	durationLap := rp.FindAllString(lapTime, -1)

	milisecondsInt, _ := strconv.Atoi(durationLap[2])
	seconds, _ := strconv.Atoi(durationLap[1])
	minutes, _ := strconv.Atoi(durationLap[0])

	return minutes, seconds, milisecondsInt
}

func (rm *PilotResultBuilder) getBestLap(allLapTime map[int]string) string {
	lapTime := []int{}

	for time := range allLapTime {
		lapTime = append(lapTime, time)
	}
	sort.Ints(lapTime)
	return allLapTime[lapTime[0]]
}
