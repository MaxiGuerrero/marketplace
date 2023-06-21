package shared

type Response struct{
	Message string `json:"message"`
}

func OK()*Response{
	return &Response{Message: "operation sucessfully"}
}

func BadRequest(message string)*Response{
	return &Response{Message: message}
}

func Unauthorized() *Response{
	return &Response{Message: "Unauthorized access"}
}

func InternalError(message string) *Response{
	return &Response{Message: message}
}