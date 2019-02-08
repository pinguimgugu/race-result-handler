package formater

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TimeFormater struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTimeFormater(t *testing.T) {
	suite.Run(t, new(TimeFormater))
}

func TestShouldBeReturnFormatedDataPassing128Seconds(t *testing.T) {
	want := "02:08"
	formated := FormatBySeconds(128)
	assert.Equal(t, want, formated)
}

func TestShouldBeReturnFormatedDataPassing50Seconds(t *testing.T) {
	want := "00:50"
	formated := FormatBySeconds(50)
	assert.Equal(t, want, formated)
}

func TestShouldBeReturnFormatedDataPassing130Seconds(t *testing.T) {
	want := "02:10"
	formated := FormatBySeconds(130)
	assert.Equal(t, want, formated)
}

func TestShouldBeReturnFormatedDataPassing60Seconds(t *testing.T) {
	want := "01:00"
	formated := FormatBySeconds(60)
	assert.Equal(t, want, formated)
}

func TestShouldBeReturnFormatedDateDiffInSeconds(t *testing.T) {
	want := "00:50"
	formated := CalculateDifferenceDates("02:00", "02:50")
	assert.Equal(t, want, formated)
}

func TestShouldBeReturnFormatedDateDiffOneMinute(t *testing.T) {
	want := "01:00"
	formated := CalculateDifferenceDates("02:00", "03:00")
	assert.Equal(t, want, formated)
}
