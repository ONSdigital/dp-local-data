package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ONSdigital/dp-local-data/config"
	"github.com/ONSdigital/dp-local-data/tasks"
	"github.com/ONSdigital/log.go/log"
	"github.com/fatih/color"
)

type task func(cfg *config.Config) error

var (
	description = "dp-local-data is a tool for cleaning CMD data out of local dev env and/or importing the prerequisite " +
		"hierarchy/codelist data required to successfully import datasets"
)

func main() {
	commands, exit := getTasks()
	if exit {
		help()
		os.Exit(0)
	}

	cfg, err := config.Get()
	if err != nil {
		logErrAndExit(err)
	}

	output("Running CMD data Dr.")
	for _, taskFunc := range commands {
		if err := taskFunc(cfg); err != nil {
			logErrAndExit(err)
		}
		fmt.Println()
	}

	output("CMD data Dr. completed successfully")
}

func getTasks() ([]task, bool) {
	doClean := flag.Bool("clean", false, "Drop all local CMD data from Neo4j and MongoDB and deletes any Zededee collections")
	doImport := flag.Bool("import", false, "Import the generic hierarchies and code lists specified in config.yml")
	flag.Bool("help", false, "Display help info")

	flag.Parse()

	commands := make([]task, 0)

	if *doClean {
		commands = append(commands, tasks.DeleteCollections, tasks.DropMongo, tasks.DropNeo4j)
	}

	if *doImport {
		commands = append(commands, tasks.BuildHierarchies, tasks.ImportCodeLists)
	}

	return commands, len(commands) == 0
}

func help() {
	fmt.Println(description)
	fmt.Println()
	fmt.Println("Usage:\n\tdp-local-data [-commands]")
	fmt.Println("Commands:")
	flag.PrintDefaults()
	fmt.Println()
}

func logErrAndExit(err error) {
	log.Event(nil, "unexpected error", log.Error(err))
	os.Exit(1)
}

func output(msg string) {
	color.Green(fmt.Sprintf("[CMD] %s", msg))
}
