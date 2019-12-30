package date

import (
	"regexp"
	"strconv"
)

func IsParseable(dateString string) bool {
	months := []string{"jan", "feb", "mar", "may", "jun", "jul", "aug", "sep", "oct", "nov", "des"}
	var definitelyMonth, definitelyYear, possibleMonth, possibleDate bool

	//splitting the date elements using regex allowed delimiter in dateString is (-) or (/) or ( )
	validDate := regexp.MustCompile(`^(\w+)[-|\/| ](\w+)[-|\/| ](\w+)$`)
	matches := validDate.FindStringSubmatch(dateString)

	//foreach matches determine is it a date, a month, or a year
	for _, match := range matches {
		//first element is always the input element, so ignore that
		if match == dateString {
			continue
		}

		//convert it to int if it can be converted,
		//if not determine is it a month in 3 letter format,
		//if not then its not parseable, return false
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
			//if not error in converting to int and length of unconverted is 4 element
			//then it's a year element
			definitelyYear = true
		} else if intMatch > 0 {
			//if it's less than 12 and definitelyMonth is not found yet its a possible month
			if intMatch < 13 && !definitelyMonth {
				//if possible month has been found then its ambiguous/ not Parseable
				if possibleMonth {
					return false
				}
				possibleMonth = true
			}

			//if it's less than 32 its a possible date
			if intMatch < 32 {
				//if possible date has been found then its ambiguous/ not Parseable
				if possibleDate && !possibleMonth && !definitelyMonth {
					return false
				}
				possibleDate = true
			}
		}
	}

	//if there is no definite year or possible date then it's not parseable
	return definitelyYear && possibleDate
}
