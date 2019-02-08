package repository

import (
	"app/domain/entity"
	"app/domain/service"
	"app/infrastructure/formater"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func (r *RaceResult) CreateClassification(classificationPilot []string, pilotStatistic map[string]entity.RacePilotStatistic) {
	fileCsv, _ := os.OpenFile("result_classification.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer fileCsv.Close()

	position := 1

	csvWriter := csv.NewWriter(fileCsv)
	strWrite := []string{"POSICAO", "CODIGO", "NOME", "VOLTAS", "TEMPO", "MEDIA VELOCIDADE CORRIDA", "MELHOR VOLTA", "DIFERENCA TEMPO PARA VENCEDOR"}
	csvWriter.Write(strWrite)

	differenceTimeToWinner := "00:00"
	timeRaceWinner := ""

	identifierBestLap := new(service.IdentifierBestLap)
	for _, pilotNumber := range classificationPilot {

		if position == 1 {
			timeRaceWinner = pilotStatistic[pilotNumber].RaceTime
		}

		if position > 1 {
			differenceTimeToWinner = formater.CalculateDifferenceDates(
				timeRaceWinner,
				pilotStatistic[pilotNumber].RaceTime,
			)
		}

		identifierBestLap.Attach(pilotStatistic[pilotNumber])
		value := []string{
			strconv.Itoa(position),
			pilotStatistic[pilotNumber].Number,
			pilotStatistic[pilotNumber].Name,
			strconv.Itoa(pilotStatistic[pilotNumber].LapAmount),
			pilotStatistic[pilotNumber].RaceTime,
			fmt.Sprintf("%f", pilotStatistic[pilotNumber].SpeedRaceAverage),
			pilotStatistic[pilotNumber].BestLap,
			differenceTimeToWinner,
		}
		fmt.Println(value)
		csvWriter.Write(value)
		csvWriter.Flush()
		position++
	}
	csvWriter.Write([]string{"MELHOR VOLTA CORRIDA"})
	csvWriter.Flush()
	csvWriter.Write([]string{"CODIGO PILOTO", "TEMPO"})
	csvWriter.Flush()

	pilotNumber, timeLap := identifierBestLap.Get()
	bestLap := []string{pilotNumber, timeLap}
	csvWriter.Write(bestLap)
	csvWriter.Flush()

	fmt.Println("Piloto | Melhor volta")
	fmt.Println(bestLap)
	fmt.Println("csv created")
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
