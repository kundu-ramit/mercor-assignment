package main

import (
	"log"
	"os"

	"github.com/kundu-ramit/mercor_assignment/cmd/skills"
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
	// db := database.Initialize()
	switch arg {
	case "fetchskills":
		//skills.FetchSkills(db)
	case "fetchskillvector":
		skills.FetchSkillVectorOpenAi()
	case "addskillvector":
		skills.AddSkillVectors()
	}

}

// func startServer() {

// 	database.Initialize()
// 	redis.Initialize()
// 	r := routes.SetupRouter()
// 	r.Run(":8002")
// }
