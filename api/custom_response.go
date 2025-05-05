package customresponse

type CustomResponse struct {
	DeveloperMessage string      `json:"developerMessage"`
	UserMessage      string      `json:"userMessage"`
	Err              error       `json:"err"`
}

func NewCustomResponse(devMessage string, userMessage string, err error) *CustomResponse {
	return &CustomResponse{
		DeveloperMessage: devMessage,
		UserMessage:      userMessage,
		Err:              err,
	}
}
