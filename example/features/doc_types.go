package features

// ErrorResponse representa a estrutura de resposta de erro
type ErrorResponse struct {
	Error string `json:"error"`
}

// HealthResponse representa a estrutura de resposta do health check
type HealthResponse struct {
	Status string `json:"status"`
}

// LoggedResponse representa a estrutura de resposta do endpoint logged
type LoggedResponse struct {
	Message string `json:"message"`
}
