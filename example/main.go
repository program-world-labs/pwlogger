package main

import "github.com/program-world-labs/pwlogger"

func main() {
	// Initialize logger
	l := pwlogger.NewDevelopmentLogger("project-id")
	// Initialize event
	l.Info().Msg("Hello World!")

	l = pwlogger.NewProductionLogger("project-id")
	l.Info().Msg("Hello World!")

}
