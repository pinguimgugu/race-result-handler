package formater

import "fmt"

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
