package model

type Request struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Smile   int    `json:"smile"`
}

type FIO struct {
	UserID    int    `json:"user_id,omitempty"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

type Letter struct {
	UserID int    `json:"user_id,omitempty"`
	Item   string `json:"item"`
	Letter string `json:"letter"`
}

type Message struct {
	UserID  int    `json:"user_id,omitempty"`
	Message string `json:"message,omitempty"`
}
