package models

// Define our message object
type Message struct {
	User    Professor `json:"user"`
	Message  string `json:"message"`
}

type Professor struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      string `json:"age"`
	Bio      string `json:"bio"`
	Email    string `json:"email"`
}
