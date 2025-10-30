package handlers

type Server struct {
	AvagencAPIKey  string
	AvagencAgenticTuyaAgentWebhookURL string
	AvagencAgenticTuyaXiaozhiAgentMCPWebsocketURL string
}

func NewServer(apiKey, AgenticTuyaAgentWebhookURL, AgenticTutaVoiceAgentMCPWebsocketURL string) *Server {
	return &Server{
		AvagencAPIKey:  apiKey,
		AvagencAgenticTuyaAgentWebhookURL: AgenticTuyaAgentWebhookURL,
		AvagencAgenticTuyaXiaozhiAgentMCPWebsocketURL: AgenticTutaVoiceAgentMCPWebsocketURL,
	}
}
