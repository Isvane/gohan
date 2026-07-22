package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
