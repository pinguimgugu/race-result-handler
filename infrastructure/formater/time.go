package formater

import "fmt"

func FormatBySeconds(timeInSeconds int) string {
	var raceTime string
	resto := timeInSeconds % 60

	if resto > 9 {
		raceTime = fmt.Sprintf("0%d:%d", timeInSeconds/60, resto)
	} else if resto == 0 {
		raceTime = fmt.Sprintf("0%d:00", timeInSeconds/60)
	} else {
		raceTime = fmt.Sprintf("0%d:0%d", timeInSeconds/60, resto)
	}

	return raceTime

}
