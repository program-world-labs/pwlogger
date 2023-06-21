package main

import logger "github.com/program-world-labs/zerolog-gcp"

func main() {
	// Initialize logger
	l := logger.NewDevelopmentLogger("project-id")
	// Initialize event
	l.Info().Msg("Hello World!")

	l = logger.NewProductionLogger("project-id")
	l.Info().Msg("Hello World!")

}
