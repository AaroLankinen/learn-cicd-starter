package main

import "net/http"

// handlerReadiness handles the health check GET /v1/healthz endpoint.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
