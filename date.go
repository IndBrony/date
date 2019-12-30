package date

import (
	"regexp"
	"strconv"
)

func IsParseable(dateString string) bool {
	months := []string{"jan", "feb", "mar", "may", "jun", "jul", "aug", "sep", "oct", "nov", "des"}
	var definitelyMonth, definitelyYear, possibleMonth, possibleDate bool
	validDate := regexp.MustCompile(`^(\w+)[-|\/| ](\w+)[-|\/| ](\w+)$`)
	matches := validDate.FindStringSubmatch(dateString)

	for _, match := range matches {
		if match == dateString {
			continue
		}
		intMatch, err := strconv.Atoi(match)
		if err != nil {
			for _, month := range months {
				if match == month {
					definitelyMonth = true
					break
				}
			}
			if !definitelyMonth {
				return false
			}
		} else if len(match) == 4 {
			definitelyYear = true
		} else if intMatch > 0 {
			if intMatch < 13 && !definitelyMonth {
				if possibleMonth {
					return false
				}
				possibleMonth = true
			}
			if intMatch < 32 {
				if possibleDate && !possibleMonth && !definitelyMonth {
					return false
				}
				possibleDate = true
			}
		}
	}
	if !definitelyYear || !possibleDate {
		return false
	}
	return true
}
