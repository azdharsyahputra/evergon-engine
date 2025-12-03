package main

import (
	"fmt"
	"os"
	"path/filepath"

	"evergon/internal/api"
	"evergon/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Base path tempat evergon berada
	base, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	engine := core.NewEngine(base)

	switch os.Args[1] {

	case "start":
		if err := engine.StartAll(); err != nil {
			fmt.Println("Error starting services:", err)
		}

	case "stop":
		if err := engine.StopAll(); err != nil {
			fmt.Println("Error stopping services:", err)
		}

	case "api":
		api.StartAPIServer(engine)

	case "php":
		handlePHPCommand(engine)

	default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
	}
}

//////////////////////////////////////////////////////
// CLI SUBCOMMAND: PHP
//////////////////////////////////////////////////////

func handlePHPCommand(engine *core.Engine) {
	if len(os.Args) < 3 {
		printPHPUsage()
		return
	}

	switch os.Args[2] {

	case "use": // evergon php use <version>
		if len(os.Args) < 4 {
			fmt.Println("Missing version.")
			printPHPUsage()
			return
		}
		ver := os.Args[3]

		if err := engine.SetPHPVersion(ver); err != nil {
			fmt.Println("Error setting PHP version:", err)
			return
		}

		fmt.Println("PHP version switched to", ver)

	case "versions": // evergon php versions
		versions, err := engine.ListPHPVersions()
		if err != nil {
			fmt.Println("Error listing versions:", err)
			return
		}
		fmt.Println("Available PHP versions:", versions)

	case "current": // evergon php current
		fmt.Println("Current PHP version:", engine.CurrentPHPVersion())

	default:
		fmt.Println("Unknown php command:", os.Args[2])
		printPHPUsage()
	}
}

//////////////////////////////////////////////////////
// HELPERS
//////////////////////////////////////////////////////

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  evergon start")
	fmt.Println("  evergon stop")
	fmt.Println("  evergon api")
	fmt.Println("  evergon php use <version>")
	fmt.Println("  evergon php versions")
	fmt.Println("  evergon php current")
}

func printPHPUsage() {
	fmt.Println("PHP Commands:")
	fmt.Println("  evergon php use <version>")
	fmt.Println("  evergon php versions")
	fmt.Println("  evergon php current")
}
