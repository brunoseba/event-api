package utils

import (
	"time"
)

func ValidateState(isAdmin bool, state string) string {
	var statetype string
	if !isAdmin {
		statetype = "publicado"

	} else {
		statetype = state
	}
	return statetype
}

func ValidateDate(dateIn, dateOut time.Time) bool {
	if dateIn.IsZero() || dateOut.IsZero() {
		return ValueFalse
	}
	return ValueTrue
}

func ValidateTitle(title string) bool {
	if title != "" {
		return ValueFalse
	}
	return ValueTrue
}

func ConvertDate(date string) time.Time {
	dateVal, _ := time.Parse(time.RFC3339, date)
	return dateVal
}
