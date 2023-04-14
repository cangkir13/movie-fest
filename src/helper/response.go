package helper

type PublicResponse struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
} //@name Global-Response
