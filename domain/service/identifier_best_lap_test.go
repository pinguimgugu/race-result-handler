package service

import (
	"app/domain/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IdentifierBestLapSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIdentifierBestLapSuite(t *testing.T) {
	suite.Run(t, new(IdentifierBestLapSuite))
}
func (suite *IdentifierBestLapSuite) TestShouldBeReturnBestLapTimeAndNumberPilotWhenHaveResutList() {
	racePilots := []entity.RacePilotStatistic{}

	wantPilot := "003"
	wantBestLap := "01:00.332"
	racePilots = append(racePilots, entity.RacePilotStatistic{Number: "001", BestLap: "01:00.333"})
	racePilots = append(racePilots, entity.RacePilotStatistic{Number: "002", BestLap: "01:01.333"})
	racePilots = append(racePilots, entity.RacePilotStatistic{Number: wantPilot, BestLap: wantBestLap})
	racePilots = append(racePilots, entity.RacePilotStatistic{Number: "004", BestLap: "02:10.333"})

	identifierBestLap := IdentifierBestLap{}
	for _, pilot := range racePilots {
		identifierBestLap.Attach(pilot)
	}
	pilotNumber, lapTime := identifierBestLap.Get()

	suite.Equal(wantPilot, pilotNumber)
	suite.Equal(wantBestLap, lapTime)

}

func (suite *IdentifierBestLapSuite) TestShouldBeReturnEmptyDataWhenHaventResutListData() {
	identifierBestLap := IdentifierBestLap{}
	pilotNumber, lapTime := identifierBestLap.Get()

	suite.Equal("", pilotNumber)
	suite.Equal("", lapTime)

}
