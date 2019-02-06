package main

import (
	"app/application"
	"app/domain/service"
	"app/infrastructure/repository"
)

func main() {

	raceResultHandler := service.NewRaceResultHandler(
		repository.NewRaceResult(),
		service.NewRacerClassifier(),
		service.NewRaceMetricBuilder(),
	)

	application.NewApplication(raceResultHandler).Run()
}
