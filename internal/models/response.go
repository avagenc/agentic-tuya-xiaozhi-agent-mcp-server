package models

type BaseResponse struct {
	Action  string `json:"action"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Timestamp int64  `json:"timestamp"`

	// Exists only if Success is true
	Result  any    `json:"result,omitempty"`

	// Exists only if Success is false
	Error   string `json:"error,omitempty"`
}
