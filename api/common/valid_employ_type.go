package common

import (
	"errors"
	"fmt"
	api_base "griffin-dao/api/base"
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

// At V2 version of validation;
// Return api_base.CommonResponse too if error.
// Clean code on the using side.

func ValidEmployTypeV2(value string) (api_base.CommonResponse, error) {
	possible := map[string]bool{
		"permanent": true,
		"freelance": true,
	}

	if !possible[value] {
		packet := api_base.CommonResponse{
			Status:  false,
			Message: fmt.Sprintf("Request wrong type. Unsupported => %s", value),
		}
		return packet, errors.New("no such employ type")
	}

	return api_base.CommonResponse{}, nil
}

func ValidPayFreqV2(value string) (api_base.CommonResponse, error) {
	// Validate whether is it "D", "W", or "M". Capitalization required
	possible := map[string]bool{
		"D": true,
		"W": true,
		"M": true,
	}
	if !possible[value] {
		packet := api_base.CommonResponse{
			Status:  false,
			Message: fmt.Sprintf("No such pay frequency => %s", value),
		}
		return packet, errors.New("no such payment frequency")
	}
	return api_base.CommonResponse{}, nil
}

func ValidMonthDateV2(day time.Time, value string) (api_base.CommonResponse, error) {
	if value == "M" && day.Day() > 25 {
		// Not all month has 29th, ~ 31th. i.e.) Feb
		packet := api_base.CommonResponse{
			Status:  false,
			Message: fmt.Sprintf("If pay frequency is month, => date %v <= is not allowed", day.Day()),
		}
		return packet, errors.New("not allowed pay date if payfreq is Month (M)")
	}

	return api_base.CommonResponse{}, nil
}
