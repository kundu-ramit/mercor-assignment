package main

import (
	"log"
	"os"

	"github.com/kundu-ramit/mercor_assignment/cmd/skills"
	"github.com/kundu-ramit/mercor_assignment/infra/database"
)

func main() {
	// Get the command-line arguments
	args := os.Args[1:]

	// Check the number of arguments
	if len(args) == 0 {
		log.Fatal("No command specified.")
	}

	// Handle the command
	switch args[0] {
	case "seed":
		applySeed(args[1])
	default:
		log.Fatal("Invalid command:", args[0])
	}
}

func applySeed(arg string) {

	switch arg {
	case "fetchskills":
		db := database.Initialize()
		//vectordb := vectordatabase.Initialize()
		skills.FetchSkills(db)
	case "fetchskillvector":
		skills.FetchSkillVectorOpenAi()
	}

}

// func startServer() {

// 	database.Initialize()
// 	redis.Initialize()
// 	r := routes.SetupRouter()
// 	r.Run(":8002")
// }
