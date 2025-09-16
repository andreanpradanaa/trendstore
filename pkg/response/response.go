package response

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
