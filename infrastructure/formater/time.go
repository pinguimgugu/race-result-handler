package formater

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FormatBySeconds(timeInSeconds int) string {
	var raceTime string
	divisionRemainder := timeInSeconds % 60

	if divisionRemainder > 9 {
		raceTime = fmt.Sprintf("0%d:%d", timeInSeconds/60, divisionRemainder)
	} else if divisionRemainder == 0 {
		raceTime = fmt.Sprintf("0%d:00", timeInSeconds/60)
	} else {
		raceTime = fmt.Sprintf("0%d:0%d", timeInSeconds/60, divisionRemainder)
	}

	return raceTime

}

func CalculateDifferenceDates(dateOne, dateTwo string) string {
	rp := regexp.MustCompile("[0-9]+")
	dateOneSlice := rp.FindAllString(dateOne, -1)
	dateOneMinute, _ := strconv.Atoi(dateOneSlice[0])
	dateOneSeconds, _ := strconv.Atoi(dateOneSlice[1])
	secondsDateOne := (dateOneMinute * 60) + dateOneSeconds

	dateTwoSlice := rp.FindAllString(dateTwo, -1)
	dateTwoMinute, _ := strconv.Atoi(dateTwoSlice[0])
	dateTwoSeconds, _ := strconv.Atoi(dateTwoSlice[1])

	secondsDateTwo := (dateTwoMinute * 60) + dateTwoSeconds
	diffInSeconds := secondsDateTwo - secondsDateOne
	return FormatBySeconds(diffInSeconds)
}

func ParseTimeToIntValue(timeFormat string) int {
	intValue, _ := strconv.Atoi(strings.Replace(strings.Replace(timeFormat, ":", "", -1), ".", "", -1))

	return intValue
}
