package common

import "errors"

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
