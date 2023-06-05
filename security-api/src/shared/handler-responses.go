package shared

type Response struct{
	Message string `json:"message"`
}

func Custom(message string)*Response{
	return &Response{Message: message}
}

func OK()*Response{
	return &Response{Message: "Operation Sucessfully"}
}