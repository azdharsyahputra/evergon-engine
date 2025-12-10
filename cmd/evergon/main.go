package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"evergon/internal/api"
	"evergon/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	base, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	engine := core.NewEngine(base)

	switch os.Args[1] {

	case "start":
		fmt.Println("[Init] Cleaning project runtimes before start...")
		engine.ForceKillAllProjectRuntimes()

		if err := engine.StartAll(); err != nil {
			fmt.Println("Error starting services:", err)
			return
		}

		go func() {
			fmt.Println("[API] Starting Evergon API on :7070 ...")
			api.StartAPIServer(engine)
		}()

		fmt.Println("Evergon fully started. API available at http://localhost:7070")
		select {} // block forever

	case "stop":
		err := engine.StopAll()
		if err != nil {
			fmt.Println("Error stopping services:", err)
		}
		os.Exit(0)

	case "api":
		fmt.Println("[Init] Cleaning project runtimes before API start...")
		engine.ForceKillAllProjectRuntimes()
		api.StartAPIServer(engine)

	case "php":
		handlePHPCommand(engine)

	case "project":
		handleProjectCommand(engine)

	default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
	}
}

////////////////////////////////////////////////////////
// PHP SUBCOMMANDS
////////////////////////////////////////////////////////

func handlePHPCommand(engine *core.Engine) {
	if len(os.Args) < 3 {
		printPHPUsage()
		return
	}

	switch os.Args[2] {

	case "use":
		if len(os.Args) < 4 {
			fmt.Println("Missing version.")
			printPHPUsage()
			return
		}
		ver := os.Args[3]

		engine.ForceKillAllProjectRuntimes()

		if err := engine.SetPHPVersion(ver); err != nil {
			fmt.Println("Error setting PHP version:", err)
			return
		}

		fmt.Println("PHP version switched to", ver)

	case "versions":
		versions, err := engine.ListPHPVersions()
		if err != nil {
			fmt.Println("Error listing versions:", err)
			return
		}
		fmt.Println("Available PHP versions:", versions)

	case "current":
		fmt.Println("Current PHP version:", engine.CurrentPHPVersion())

	default:
		fmt.Println("Unknown php command:", os.Args[2])
		printPHPUsage()
	}
}

////////////////////////////////////////////////////////
// PROJECT SUBCOMMANDS (FINAL)
////////////////////////////////////////////////////////

func handleProjectCommand(engine *core.Engine) {
	if len(os.Args) < 3 {
		printProjectUsage()
		return
	}

	reg := core.NewProjectRegistry(engine.BasePath)

	switch os.Args[2] {

	case "list":
		projects, err := reg.List()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, p := range projects {
			fmt.Println("-", p)
		}

	case "info":
		if len(os.Args) < 4 {
			fmt.Println("Missing project name.")
			return
		}
		name := os.Args[3]

		cfg, err := reg.LoadConfig(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Project:", cfg.Name)
		fmt.Println("PHP Version:", cfg.PHPVersion)
		fmt.Println("Port:", cfg.Port)
		fmt.Println("Root:", cfg.Root)

	case "set-port":
		if len(os.Args) < 5 {
			fmt.Println("Usage: evergon project set-port <name> <port>")
			return
		}

		name := os.Args[3]
		portStr := os.Args[4]

		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("Invalid port:", portStr)
			return
		}

		cfg, err := reg.LoadConfig(name)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		cfg.Port = port

		if err := reg.SaveConfig(name, cfg); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

		fmt.Println("Restarting project", name)

		peng, err := reg.Load(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		_ = peng.Stop()
		if err := peng.Start(); err != nil {
			fmt.Println("Failed to restart:", err)
			return
		}

		fmt.Println("Project", name, "updated to port", port)

	case "restart":
		if len(os.Args) < 4 {
			fmt.Println("Missing project name.")
			return
		}
		name := os.Args[3]

		peng, err := reg.Load(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		_ = peng.Stop()
		if err := peng.Start(); err != nil {
			fmt.Println("Restart failed:", err)
			return
		}

		fmt.Println("Project restarted:", name)

	default:
		fmt.Println("Unknown project command:", os.Args[2])
		printProjectUsage()
	}
}

////////////////////////////////////////////////////////
// HELPERS
////////////////////////////////////////////////////////

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  evergon start")
	fmt.Println("  evergon stop")
	fmt.Println("  evergon api")
	fmt.Println("  evergon php use <version>")
	fmt.Println("  evergon php versions")
	fmt.Println("  evergon php current")
	fmt.Println("  evergon project list")
	fmt.Println("  evergon project info <name>")
	fmt.Println("  evergon project set-port <name> <port>")
	fmt.Println("  evergon project restart <name>")
}

func printPHPUsage() {
	fmt.Println("PHP Commands:")
	fmt.Println("  evergon php use <version>")
	fmt.Println("  evergon php versions")
	fmt.Println("  evergon php current")
}

func printProjectUsage() {
	fmt.Println("Project Commands:")
	fmt.Println("  evergon project list")
	fmt.Println("  evergon project info <name>")
	fmt.Println("  evergon project set-port <name> <port>")
	fmt.Println("  evergon project restart <name>")
}
