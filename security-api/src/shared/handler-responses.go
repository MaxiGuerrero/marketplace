package shared

// Define how to send response messages in a request, this struct is use along of the system.
type Response struct{
	Message string `json:"message"`
}

// Default response when a operation is success.
func OK()*Response{
	return &Response{Message: "Successful operation"}
}

// Response when a operation is has a bad request. Is necessary to set a message.
func BadRequest(message string)*Response{
	return &Response{Message: message}
}

// Default response when an user doesn't have permission to use an endpoint.
func Unauthorized() *Response{
	return &Response{Message: "Unauthorized access"}
}

// Resoponse when exist an internal error. Is necessary to set a message.
func InternalError(message string) *Response{
	return &Response{Message: message}
}

func TokenValidated()*Response{
	return &Response{Message: "Token is correct"}
}