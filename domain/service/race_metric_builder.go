package service

import (
	"app/domain/entity"
	"app/infrastructure/formater"

	"math"
	"regexp"
	"strconv"
	"strings"
)

type RaceMetricBuilder struct {
}

func NewRaceMetricBuilder() *RaceMetricBuilder {
	return &RaceMetricBuilder{}
}

func (rm *RaceMetricBuilder) Build(resultList []entity.RacePilotStatistic) entity.RacePilotStatistic {
	var miliseconds int
	var seconds int
	var minutes int
	var totalSecondsDuration int
	var speedAverage float64
	var speedAmount float64

	pilotStatistic := entity.RacePilotStatistic{
		LapAmount: len(resultList),
	}

	for _, data := range resultList {
		pilotStatistic.Name = data.Name
		pilotStatistic.Number = data.Number

		rp := regexp.MustCompile("[0-9]+")
		durationLap := rp.FindAllString(data.LapTime, -1)

		milisecondsInt, _ := strconv.Atoi(durationLap[2])
		seconds, _ = strconv.Atoi(durationLap[1])
		minutes, _ = strconv.Atoi(durationLap[0])

		f, _ := strconv.ParseFloat(
			strings.Replace(data.SpeedLapAverage, ",", ".", -1), 64,
		)
		totalSecondsDuration += (seconds + 60*minutes)
		speedAmount += f
		miliseconds += milisecondsInt
	}

	speedAverage = speedAmount / float64(len(resultList))
	totalSecondsDuration += int(math.Ceil(float64(miliseconds) / float64(1000)))

	pilotStatistic.SpeedRaceAverage = speedAverage
	pilotStatistic.RaceTime = formater.FormatBySeconds(totalSecondsDuration)
	return pilotStatistic
}
