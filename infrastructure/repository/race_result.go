package repository

import (
	"app/domain/entity"
	"bufio"
	"os"
	"regexp"
)

type RaceResult struct {
	raceList []entity.RacePilotStatistic
}

func NewRaceResult() *RaceResult {
	return &RaceResult{make([]entity.RacePilotStatistic, 0)}
}

func (r *RaceResult) GetList() []entity.RacePilotStatistic {
	file, _ := os.Open("/go/src/app/race_log.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		if r.isTopOfFile(line) {
			line++
			continue
		}
		r.raceList = append(r.raceList, r.parseLine(scanner.Text()))
	}
	return r.raceList
}

func (r *RaceResult) isTopOfFile(line int) bool {
	return line == 0
}

func (r *RaceResult) parseLine(line string) entity.RacePilotStatistic {
	pilotStatistic := entity.RacePilotStatistic{}
	rp := regexp.MustCompile("[a-zA-Z.0-9:,]+")
	parsedFile := rp.FindAllString(line, -1)

	pilotStatistic.Number = parsedFile[1]
	pilotStatistic.Name = parsedFile[2]
	pilotStatistic.Lap = parsedFile[3]
	pilotStatistic.LapTime = parsedFile[4]
	pilotStatistic.SpeedLapAverage = parsedFile[5]

	return pilotStatistic
}
