package service

import (
	"app/domain/contract"
	"app/domain/entity"
)

type RaceResultHandler struct {
	raceResultRepository contract.RaceResultRepository
	racerClassifier      contract.RacerClassifier
	PilotResultBuilder   contract.PilotResultBuilder
}

func NewRaceResultHandler(rrr contract.RaceResultRepository, rc contract.RacerClassifier, rmb contract.PilotResultBuilder) *RaceResultHandler {
	return &RaceResultHandler{rrr, rc, rmb}
}

func (rh *RaceResultHandler) CreateDetailed() {
	groupedResultByPilot := rh.groupResultByPilot(
		rh.raceResultRepository.GetList(),
	)

	pilotMetricChan := make(chan map[string]entity.RacePilotStatistic)
	raceClassificationChan := make(chan []string)

	go func() {
		pilotMetric := make(map[string]entity.RacePilotStatistic)
		for pilotNumber, laps := range groupedResultByPilot {
			pilotMetric[pilotNumber] = rh.PilotResultBuilder.Build(laps)
		}
		pilotMetricChan <- pilotMetric
	}()

	go func() {
		raceClassificationChan <- rh.racerClassifier.Make(groupedResultByPilot)
	}()

	rh.raceResultRepository.CreateClassification(<-raceClassificationChan, <-pilotMetricChan)
}

func (rh *RaceResultHandler) groupResultByPilot(resultList []entity.RacePilotStatistic) map[string][]entity.RacePilotStatistic {
	groupedResult := map[string][]entity.RacePilotStatistic{}

	for _, data := range resultList {
		groupedResult[data.Number] = append(groupedResult[data.Number], data)
	}

	return groupedResult
}
