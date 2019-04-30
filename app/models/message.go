package models

// Message perte responder al cliente
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
