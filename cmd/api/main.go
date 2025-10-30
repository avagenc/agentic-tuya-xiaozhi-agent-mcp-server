package main

import (
	"log"
	"net/http"
	"os"

	"github.com/avagenc/agentic-tuya-xiaozhi-agent-mcp-server/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Reading credentials from environment.")
	}

	log.Println("--- Reading Environment Variables ---")

	avagencAPIKey := os.Getenv("AVAGENC_API_KEY")
	log.Printf("DEBUG: AVAGENC_API_KEY = %s", avagencAPIKey)
	if avagencAPIKey == "" {
		log.Fatal("FATAL: AVAGENC_API_KEY environment variable not set. Service cannot start.")
	}

	avagencAgenticTuyaAgentWebhookURL := os.Getenv("AVAGENC_AGENTIC_TUYA_AGENT_WEBHOOK_URL")
	log.Printf("DEBUG: AVAGENC_AGENTIC_TUYA_AGENT_WEBHOOK_URL = %s", avagencAgenticTuyaAgentWebhookURL)

	avagencAgenticTuyaXiaozhiAgentMCPWebsocketURL := os.Getenv("AVAGENC_AGENTIC_TUYA_XIAOZHI_AGENT_MCP_WEBSOCKET_URL")
	log.Printf("DEBUG: AVAGENC_AGENTIC_TUYA_XIAOZHI_AGENT_MCP_WEBSOCKET_URL = %s", avagencAgenticTuyaXiaozhiAgentMCPWebsocketURL)


	log.Println("--- Finished Reading Environment Variables ---")

	if avagencAgenticTuyaAgentWebhookURL == "" {
		log.Fatal("Agent Webhoook URL is not set")
	}

	log.Println("DEBUG: All required environment variables are present.")
	log.Println("DEBUG: Initializing server...")

	server := handlers.NewServer(avagencAPIKey, avagencAgenticTuyaAgentWebhookURL, avagencAgenticTuyaXiaozhiAgentMCPWebsocketURL)

	http.HandleFunc("/", server.RootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("DEBUG: Attempting to listen on port %s", port)

	log.Printf("Starting Avagenc Agentic Tuya Xiaozhi to Agent Bridge on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}

func authMiddleware(apiKey string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientAPIKey := r.Header.Get("x-avagenc-api-key")
		if clientAPIKey != apiKey {
			log.Printf("Authentication failed: Invalid API Key. Request from %s", r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
