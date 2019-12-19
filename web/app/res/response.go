package res

type ValidationError struct {
	Message   string `json:"message" example:"should be a integer"`
	FieldName string `json:"field_name" example:"subject_id"`
}

type Json struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Message string `json:"message" example:"got a error in server "`
	Status  string `json:"status" example:"error"`
}
