package res

type ValidationError struct {
	Message   string `json:"message" example:"should be a integer"`
	FieldName string `json:"field_name" example:"subject_id"`
}

type Error struct {
	Message string `json:"message" example:"record not found/missing input/got a error in server"`
	Status  string `json:"status" example:"error"`
}
