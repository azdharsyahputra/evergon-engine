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
		fmt.Println("Usage:")
		fmt.Println("  evergon start")
		fmt.Println("  evergon stop")
		fmt.Println("  evergon api")
		return
	}

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
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
