package types

type ResponseError struct {
	Err Error `json:"error"`
}

func (responseError *ResponseError) Error() string {
	return responseError.Err.Message
}

type Error struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}
