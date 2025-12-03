package api

import (
	"fmt"
	"net/http"

	"evergon/internal/core"
)

func StartAPIServer(engine *core.Engine) {
	mux := http.NewServeMux()

	// register engine routes
	RegisterRoutes(mux, engine)

	// NEW: project routes
	projectHandler := NewProjectHandler(engine.BasePath)
	projectHandler.Register(mux)

	fmt.Println("Evergon API running at http://localhost:7070")
	http.ListenAndServe(":7070", mux)
}
