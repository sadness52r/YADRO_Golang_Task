package customresponse

type CustomResponse struct {
	DeveloperMessage string      `json:"developerMessage"`
	UserMessage      string      `json:"userMessage"`
	Data             interface{} `json:"data"`
	Err              error       `json:"err"`
}

func NewCustomResponse(devMessage string, userMessage string, data interface{}, err error) *CustomResponse {
	return &CustomResponse{
		DeveloperMessage: devMessage,
		UserMessage:      userMessage,
		Data:             data,
		Err:              err,
	}
}
