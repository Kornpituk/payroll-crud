package api

type HelloRequest struct {
}

type HelloResponse struct {
	Message string `json:"message"`
}
