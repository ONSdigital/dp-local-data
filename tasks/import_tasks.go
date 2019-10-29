package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ONSdigital/dp-local-data/config"
	"github.com/fatih/color"
)

var (
	goPath               = os.Getenv("GOPATH")
	onsDigitalPath       = filepath.Join(goPath, "src/github.com/ONSdigital")
	hierarchyBuilderPath = filepath.Join(onsDigitalPath, "dp-hierarchy-builder/cypher-scripts")
	codeListScriptsPath  = filepath.Join(onsDigitalPath, "dp-code-list-scripts/code-list-scripts")
)

func BuildHierarchies(cfg *config.Config) error {
	if len(cfg.Hierarchies) == 0 {
		output("No hierarchies defined in config skipping step")
		return nil
	}

	output(fmt.Sprintf("Building generic hierarchies: %+v", cfg.Hierarchies))

	for _, script := range cfg.Hierarchies {
		command := fmt.Sprintf("cypher-shell < %s/%s", hierarchyBuilderPath, script)

		if err := execCommand(command, ""); err != nil {
			return err
		}
	}

	output("Hierarchies built successfully")
	return nil
}

func ImportCodeLists(cfg *config.Config) error {
	if len(cfg.Codelists) == 0 {
		output("No code lists defined in config skipping step")
		return nil
	}

	output(fmt.Sprintf("Importing code lists: %+v", cfg.Codelists))

	for _, codelist := range cfg.Codelists {
		command := fmt.Sprintf("./load -q=%s -f=%s", "cypher", codelist)

		if err := execCommand(command, codeListScriptsPath); err != nil {
			return err
		}
	}

	output("Code lists imported successfully")
	return nil
}

func execCommand(command string, wrkDir string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr

	if len(wrkDir) > 0 {
		cmd.Dir = wrkDir
	}

	return cmd.Run()
}

func output(msg string) {
	color.Magenta(fmt.Sprintf("[import] %s", msg))
}
