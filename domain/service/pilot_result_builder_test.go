package service

import (
	"app/domain/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PilotResultBuilderSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPilotResultBuilderSuite(t *testing.T) {
	suite.Run(t, new(PilotResultBuilderSuite))
}
func (suite *PilotResultBuilderSuite) TestShouldBeReturnRacePilotStatisticBestLapWithLessTimeLapOfPilot() {
	racesPilot := []entity.RacePilotStatistic{}

	pilot := entity.RacePilotStatistic{Number: "001", Name: "test", SpeedLapAverage: "44.444"}
	pilot.LapTime = "01:01.333"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:01.000"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:03.900"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:00.200"
	racesPilot = append(racesPilot, pilot)

	pilotResultBuilder := NewPilotResultBuilder()
	pilotStatistic := pilotResultBuilder.Build(racesPilot)
	suite.Equal("01:00.200", pilotStatistic.BestLap)
}

func (suite *PilotResultBuilderSuite) TestShouldBeHaveRaceTimeWithSumAllTimeLapWhitoutMiliseconds() {
	racesPilot := []entity.RacePilotStatistic{}

	pilot := entity.RacePilotStatistic{Number: "001", Name: "test", SpeedLapAverage: "44.444"}
	pilot.LapTime = "01:00.000"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:00.000"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:00.000"
	racesPilot = append(racesPilot, pilot)

	pilotResultBuilder := NewPilotResultBuilder()
	pilotStatistic := pilotResultBuilder.Build(racesPilot)
	suite.Equal("03:00", pilotStatistic.RaceTime)
}

func (suite *PilotResultBuilderSuite) TestShouldBeHaveRaceTimeWithSumAllTimeLapWithMiliseconds() {
	racesPilot := []entity.RacePilotStatistic{}

	pilot := entity.RacePilotStatistic{Number: "001", Name: "test", SpeedLapAverage: "44.444"}
	pilot.LapTime = "01:00.000"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:00.200"
	racesPilot = append(racesPilot, pilot)
	pilot.LapTime = "01:00.300"
	racesPilot = append(racesPilot, pilot)

	pilotResultBuilder := NewPilotResultBuilder()
	pilotStatistic := pilotResultBuilder.Build(racesPilot)
	suite.Equal("03:01", pilotStatistic.RaceTime)
}

func (suite *PilotResultBuilderSuite) TestShouldBeHaveRightSpeedRaceAverageByLaps() {
	racesPilot := []entity.RacePilotStatistic{}

	pilot := entity.RacePilotStatistic{Number: "001", Name: "test", LapTime: "01:00.000"}
	pilot.SpeedLapAverage = "10,0"
	racesPilot = append(racesPilot, pilot)
	pilot.SpeedLapAverage = "10,0"
	racesPilot = append(racesPilot, pilot)
	pilot.SpeedLapAverage = "10,0"
	racesPilot = append(racesPilot, pilot)

	pilotResultBuilder := NewPilotResultBuilder()
	pilotStatistic := pilotResultBuilder.Build(racesPilot)
	suite.Equal(10.0, pilotStatistic.SpeedRaceAverage)
}
