package service

import (
	"errors"
	"fmt"
)

func ValidateContractPeriod(period int, empType string) error {
	switch {
	case period <= 0 && empType != "permanent":
		return errors.New(fmt.Sprintf("non-positive period %v, empType should be permanent", period))
	case period > 0 && empType == "permanent":
		return errors.New(fmt.Sprintf("positive work period %v, empType should be contract", period))
	default:
		return nil
	}
}
