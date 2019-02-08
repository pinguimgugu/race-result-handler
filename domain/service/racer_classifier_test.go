package service

import (
	"app/domain/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RacerClassifierSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRacerClassifierSuite(t *testing.T) {
	suite.Run(t, new(RacerClassifierSuite))
}
func (suite *RacerClassifierSuite) TestShouldBeReturnSortedListWithPiloNumberByListLapsEachPilot() {
	groupedResult := map[string][]entity.RacePilotStatistic{}

	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:02.000"})
	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:01.700"})
	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:06.000"})

	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:02.000"})
	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:01.000"})
	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:06.000"})

	groupedResult["003"] = append(groupedResult["003"], entity.RacePilotStatistic{Number: "003", LapTime: "1:30.000"})
	groupedResult["003"] = append(groupedResult["003"], entity.RacePilotStatistic{Number: "003", LapTime: "1:30.000"})
	groupedResult["003"] = append(groupedResult["003"], entity.RacePilotStatistic{Number: "003", LapTime: "1:30.000"})
	racerClassifier := NewRacerClassifier()
	classificationRace := racerClassifier.Make(groupedResult)
	suite.Equal([]string{"002", "001", "003"}, classificationRace)
}

func (suite *RacerClassifierSuite) TestShouldBeReturnSortedListWithPiloNumberAndLastPositionPilotWithTwoCompletedLaps() {
	groupedResult := map[string][]entity.RacePilotStatistic{}

	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:02.000"})
	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:01.700"})
	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:06.000"})

	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:02.000"})
	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:01.000"})
	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:06.000"})

	groupedResult["003"] = append(groupedResult["003"], entity.RacePilotStatistic{Number: "003", LapTime: "1:30.000"})
	groupedResult["003"] = append(groupedResult["003"], entity.RacePilotStatistic{Number: "003", LapTime: "1:30.000"})
	racerClassifier := NewRacerClassifier()
	classificationRace := racerClassifier.Make(groupedResult)
	suite.Equal([]string{"002", "001", "003"}, classificationRace)
}

func (suite *RacerClassifierSuite) TestShouldBeReturnSortedListWithPiloNumberAndLastPositionPilotWithOneCompletedLaps() {
	groupedResult := map[string][]entity.RacePilotStatistic{}

	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:02.000"})
	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:01.700"})
	groupedResult["001"] = append(groupedResult["001"], entity.RacePilotStatistic{Number: "001", LapTime: "1:06.000"})

	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:02.000"})
	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:01.000"})
	groupedResult["002"] = append(groupedResult["002"], entity.RacePilotStatistic{Number: "002", LapTime: "1:06.000"})

	groupedResult["003"] = append(groupedResult["003"], entity.RacePilotStatistic{Number: "003", LapTime: "1:30.000"})
	racerClassifier := NewRacerClassifier()
	classificationRace := racerClassifier.Make(groupedResult)
	suite.Equal([]string{"002", "001", "003"}, classificationRace)
}
