package reponse

type Reponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(message string, data interface{}) Reponse {
	return Reponse{
		Message: message,
		Data:    data,
	}
}
