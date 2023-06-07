package shared

type Response struct{
	Message string `json:"message"`
}

func OK()*Response{
	return &Response{Message: "Operation Sucessfully"}
}

func BadRequest(message string)*Response{
	return &Response{Message: message}
}