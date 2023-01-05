package common

import (
	"errors"
	"time"
)

func ValidateEmployType(value string) error {
	possible := map[string]bool{
		"permanent": true,
		"freelance": true,
	}
	if !possible[value] {
		return errors.New("no such employ type")
	}
	return nil
}

func ValidatePayFreq(value string) error {
	possible := map[string]bool{
		"D": true,
		"W": true,
		"M": true,
	}
	if !possible[value] {
		return errors.New("no such pay frequency")
	}
	return nil
}

func ValidateMonthDate(day time.Time, value string) error {
	if value != "M" {
		return nil
	}
	if day.Day() > 25 {
		return errors.New("payday with interval `M` should not be later than 25th")
	}
	return nil
}
