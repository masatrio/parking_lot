package cmd

import (
	"errors"

	"github.com/masatrio/parking_lot/handler"
	"github.com/masatrio/parking_lot/util"
)

type validator struct{}

func Validator() handler.Validator {
	return &validator{}
}

func (v *validator) Validate(args []string) error {
	if len(args) < 1 {
		return errors.New(util.Err.NotEnoughArgument)
	}

	switch args[0] {
	case "leave",
		"registration_numbers_for_cars_with_colour",
		"slot_numbers_for_cars_with_colour",
		"create_parking_lot",
		"slot_number_for_registration_number":
		{
			if len(args) < 2 {
				return errors.New(util.Err.NotEnoughArgument)
			}
		}
	case "park":
		{
			if len(args) < 3 {
				return errors.New(util.Err.NotEnoughArgument)
			}
		}

	default:
		return errors.New("Invalid Command")
	}

	return nil

}
