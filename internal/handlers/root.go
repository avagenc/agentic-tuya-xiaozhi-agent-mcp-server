package handlers

import "net/http"

func (server *Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	const action = "Health Check"

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", action)
		return
	}

	healthCheckPayload := struct {
		Service string `json:"service"`
		Version string `json:"version"`
		Status  string `json:"status"`
	}{
		Service: "agentic-tuya-xiaozhi-to-agent-bridge",
		Version: "0.1.0",
		Status:  "ok",
	}

	writeSuccessResponse(w, http.StatusOK, healthCheckPayload, action)
}
