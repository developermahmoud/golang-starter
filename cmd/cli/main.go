package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/fatih/color"
)

func main() {
	// Load env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version: 1.0.0")
	case "migrate":
		migrateDB()
	case "seed":
		model := arg2
		if model == "" {
			color.Red("Please enter model name")
			os.Exit(0)
		}
		seedByModel(model)
	case "module":
		module := arg2
		if module == "" {
			color.Red("Please enter module name")
			os.Exit(0)
		}
		createModule(module)
	default:
		log.Println(arg2, arg3)
	}
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: command required!")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available commands:
	help    - show the help commands
	version - print application version
	`)
}

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}
	os.Exit(0)
}
