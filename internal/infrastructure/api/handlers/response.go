package handlers // ou response, depende de onde você colocar

// ErrorResponse representa uma resposta de erro da API
type ErrorResponse struct {
	Message string `json:"message"`
}
