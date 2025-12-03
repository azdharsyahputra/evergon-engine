package api

import (
	"evergon/internal/core"
	"fmt"
	"net/http"
)

func StartAPIServer(engine *core.Engine) {
	mux := http.NewServeMux()

	RegisterRoutes(mux, engine)

	fmt.Println("Evergon API running at http://localhost:7070")
	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		fmt.Println("API Server Error:", err)
	}
}
