package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// NOT expecting a command name
// if string len = 1, just try to run that, otherwise send to be parsed
func readSubcommandInput(input string) ([]string, error) {
	if input == "" {
		fmt.Println("No input found")
		return []string{}, nil
	}

	splitArgs := strings.Split(input, " ")
	for i, arg := range splitArgs {
		splitArgs[i] = strings.TrimSpace(arg)
	}

	return splitArgs, nil

}

// maps is already reference type
// used in commands
// DONT HANDLE QUOTES
// JUST REPLACE _ WITH SPACES
// BINGOBANGO
func parseArguments(flags map[string]Flag, parameters []string) ([]Argument, error) {
	// used when in a subcommand, not expecting a command name. just give it the subcommand map.
	// flags -p, command like things do not so figure out how to do that
	// flags should also take a value for now so exploit that
	if len(parameters) == 0 {
		return nil, errors.New("nothing entered. Please try again")
	}

	// split breaks on spaces, for when entering a value with a space like first last names
	// can use an underscore that will be replaced before added to DB
	underscore := false
	for _, param := range parameters {
		if strings.Contains(param, "_") {
			underscore = true
		}
	}

	var arguments []Argument

	for i := 0; i < len(parameters); i++ {
		// RIGHT NOW we're looking for "-" as a contains, but we need the first char only
		// how to we golang the first char of a string

		flag, ok := flags[parameters[i]]
		if !ok {
			err := fmt.Sprintf("%s is not a flag allowed for this command", parameters[i])
			return nil, errors.New(err)
		}

		tArg := Argument{}
		if flag.takesValue {
			tArg.flag = parameters[i]
			if i+1 == len(parameters) || string(parameters[i+1][0]) == "-" {
				err := fmt.Sprintf("%s is a flag that takes a value", parameters[i])
				return nil, errors.New(err)
			}
			i++
			tArg.value = parameters[i]
		} else {
			tArg.flag = parameters[i]
		}
		arguments = append(arguments, tArg)

	}

	// can't use range because it works via value and not reference, wont copy changes
	if underscore {
		for i := 0; i < len(arguments); i++ {
			arguments[i].value = strings.Replace(arguments[i].value, "_", " ", -1)
		}
	}

	return arguments, nil
}

func parseDate(input string) (time.Time, error) {
	// create an array of the formats (with 0s, without, 4 digit year, 2 digit year)
	// go through parse works and then return
	var date time.Time
	var err error
	timeFormats := []string{"1/2/06", "1/2/2006", "01/02/06", "01/02/2006"}
	for _, format := range timeFormats {
		date, err = time.Parse(format, input)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Println("Error parsing date.")
		return time.Time{}, err
	}

	return date, nil
}

// Returns midnight of input
// TODO: what time are my current stamps in?
func normalizeDate(t time.Time) time.Time {
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return midnight
}
