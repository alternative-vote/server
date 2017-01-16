package generated

type APIError struct {
	StatusCode       int      `json:"statusCode"`
	ErrorMessage     string   `json:"message,omitempty"`
	ErrorDetails     string   `json:"details,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
}

func HttpError(statusCode int) *APIError {
	ret := new(APIError)
	ret.StatusCode = statusCode
	var defaultMessage string
	switch statusCode {
	case 400:
		defaultMessage = "Bad Request"
	case 401:
		defaultMessage = "Unauthorized"
	case 403:
		defaultMessage = "Forbidden"
	case 404:
		defaultMessage = "Not Found"
	case 405:
		defaultMessage = "Method Not Allowed"
	case 418:
		defaultMessage = "I am a teapot."
	case 429:
		defaultMessage = "Too Many Requests"
	case 500:
		defaultMessage = "Internal Service Error"
	case 501:
		defaultMessage = "Not Implemented"
	}
	ret.ErrorMessage = defaultMessage
	return ret
}

func (o *APIError) Message(message string) *APIError {
	ret := new(APIError)
	ret.StatusCode = o.StatusCode
	ret.ErrorMessage = message
	ret.ErrorDetails = o.ErrorDetails
	return ret
}

func (o *APIError) Details(details string) *APIError {
	ret := new(APIError)
	ret.StatusCode = o.StatusCode
	ret.ErrorMessage = o.ErrorMessage
	ret.ErrorDetails = details
	return ret
}
