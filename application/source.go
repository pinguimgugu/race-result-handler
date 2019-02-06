package application

import (
	"app/domain/contract"
)

type Application struct {
	raceResultHandler contract.RaceResultHandler
}

func NewApplication(rrh contract.RaceResultHandler) *Application {
	return &Application{rrh}
}

func (a *Application) Run() {
	a.raceResultHandler.CreateDetailed()
}
