package res

type Error struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"got a error because ..."`
}
