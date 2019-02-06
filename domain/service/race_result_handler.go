package service

import (
	"app/domain/contract"
	"app/domain/entity"
	"fmt"
	"sync"
)

type RaceResultHandler struct {
	raceResultRepository contract.RaceResultRepository
	racerClassifier      contract.RacerClassifier
	raceMetricBuilder    contract.RaceMetricBuilder
}

func NewRaceResultHandler(rrr contract.RaceResultRepository, rc contract.RacerClassifier, rmb contract.RaceMetricBuilder) *RaceResultHandler {
	return &RaceResultHandler{rrr, rc, rmb}
}

func (rh *RaceResultHandler) CreateDetailed() {
	groupedResultByPilot := rh.groupResultByPilot(
		rh.raceResultRepository.GetList(),
	)

	var wg sync.WaitGroup
	wg.Add(2)

	pilotMetricChan := make(chan map[string]entity.RacePilotStatistic)
	raceClassificationChan := make(chan []string)

	go func() {
		pilotMetric := make(map[string]entity.RacePilotStatistic)
		for pilotNumber, laps := range groupedResultByPilot {
			pilotMetric[pilotNumber] = rh.raceMetricBuilder.Build(laps)
		}

		pilotMetricChan <- pilotMetric
		wg.Done()
	}()

	go func() {
		raceClassificationChan <- rh.racerClassifier.Make(groupedResultByPilot)
		wg.Done()
	}()

	go func() {
		wg.Wait()
	}()

	pilotMetricList := <-pilotMetricChan
	for _, data := range <-raceClassificationChan {
		fmt.Println(pilotMetricList[data])
	}
}

func (rh *RaceResultHandler) groupResultByPilot(resultList []entity.RacePilotStatistic) map[string][]entity.RacePilotStatistic {
	groupedResult := map[string][]entity.RacePilotStatistic{}

	for _, data := range resultList {
		groupedResult[data.Number] = append(groupedResult[data.Number], data)
	}

	return groupedResult
}
