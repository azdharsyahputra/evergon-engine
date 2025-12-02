package api

import (
	"fmt"
	"net/http"

	"evergon/internal/core"
)

func StartAPIServer(engine *core.Engine) {
	mux := http.NewServeMux()
	RegisterRoutes(mux, engine)

	fmt.Println("Evergon API running at http://localhost:9091")
	if err := http.ListenAndServe(":9091", mux); err != nil {
		fmt.Println("API server error:", err)
	}
}
