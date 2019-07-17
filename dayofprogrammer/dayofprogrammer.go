package dayofprogrammer

import "fmt"
import "time"

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	var isLeap bool
	if year == 1918 {
		start := time.Date(int(year), 2, 14, 0, 0, 0, 0, time.UTC)
		//February 14th is the 32nd day of the year in 1918
		current := start.AddDate(0, 0, 256-32)
		return current.Format("02.01.2006")
	}

	if year < 1918 {
		isLeap = year%4 == 0
	} else {
		isLeap = (year%400 == 0) || ((year%4 == 0) && (year%100 != 0))
	}

	if isLeap {
		return fmt.Sprintf("12.09.%d", year)
	}

	return fmt.Sprintf("13.09.%d", year)
}
