package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jsandberg07/clitest/internal/database"

	_ "github.com/lib/pq"
)

// COME BACK TO IT:
// writing tests. spin up a fake db and do tests there. everything is too tight
// parsing works and that's like the one thing that isn't tied to a db
// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
// https://circleci.com/blog/unit-testing-vs-integration-testing/
// use cmp, or reflect. I guess cmp is better. you can get the difference between tho values

// australia
// china #9
// china #11
// do it

// Next:
// logins (crpyto, storing that, creating new people, the admin tier account)
// so we just always login as the same thing automatically but for 'production' and later
// 1. passwords and hashing. use a different package. compare and hash. ect. do like 3 tries before exiting.
// 2. login interface
// 3. option to load test data and storing the fact that test data was stored
// 4. the option for anybody being able to activate anybodys cards that is already there and can be toggled
// 5. uhh create position. "User will be prompted to add password on first login"
// 6. then prompt password, exit, then ask them to login
// 7. Then always ask for it after
// need an investigator update password,
// get hash where $=id

// After that:
// work on the read me cause it'll be different

// AFTER THAT:
// the great polishing
// have the 'or enter x to exit' return const string 'exit' or whatever, check for that instead of like
// 0 value struct and exit from that aka the great exiting
// do you want to load test data? and allow for that'
// add more test data
// allow reset
// ask each time you login if you want to keep using test data if loaded (store that in settings)
// then we done! could easily be done by the end of the day and then make a list of what esle to do

// AFTER THAT:
// the great readme-en-ing. write a readme and update the github page
// add some consts for the strings you use for consistency + fancy + easier changes
// like exiting program vs exiting command
// AFTER THAT:
// the great adding the project to my portfolioening
// lots of test data! for fun!
// AFTER THAT:
// the great job applyening, apply for jobs
// DURING THAT:
// the adding automated DB testing as well! cause that's a thing!
// ADDITIONALLY:
// a RESTful server version! throw out all your cli code and keep the sql! no more parsing! just json!

// change if you want a million things printed or not
const verbose bool = false

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load env file: %s\n", err)
		os.Exit(1)
	}
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Could not open DB: %s\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	fmt.Println("Hello borld")
	cfg := Config{
		currentState:         nil,
		nextState:            getMainState(),
		db:                   dbQueries,
		loggedInInvestigator: nil,
		loggedInPosition:     nil,
	}

	err = cfg.db.ResetDatabase(context.Background())
	if err != nil {
		fmt.Printf("Error resetting DB: %s", err)
		os.Exit(1)
	}

	err = cfg.loadSettings()
	if err != nil {
		fmt.Printf("Error checking settings from DB: %s", err)
		os.Exit(1)
	}

	err = cfg.testData()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)

	err = cfg.login()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg.printLogin()
	fmt.Println("\n* Welcome to Musmus!")

	err = getTodaysReminders(&cfg)
	if err != nil {
		fmt.Println("Error getting today's reminders")
		fmt.Println(err)
	}

	err = getTodaysOrders(&cfg)
	if err != nil {
		fmt.Println("Error getting today's orders")
		fmt.Println(err)
	}

	// spacing :^3
	fmt.Println("")

	for {
		// check if new state
		if cfg.nextState != nil {
			cfg.currentState = cfg.nextState
			cfg.nextState = nil
		}

		fmt.Printf(">%s - ", cfg.currentState.cliMessage)

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input string: %s", err)
			os.Exit(1)
		}

		cmdName, err := readCommandName(text)
		if err != nil {
			fmt.Println("oops error getting command name")
			fmt.Println(err)
			os.Exit(1)
		}

		command, ok := cfg.currentState.currentCommands[cmdName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		/* removed because arguments are no longer passed to commands (they were just never used)
		// check to see if the flags are available, and if they take values, return flags and args
		arguments, err := parseCommandArguments(&command, parameters)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*/

		// pass the args into the commands function, then run it
		err = command.function(&cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// spacing :^3
		fmt.Println()
	}

	// cool facts: this part of the code is never reached beacuse exit uses os dot Exit(0)
}
