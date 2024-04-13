package broker

type Response struct {
	Data      any      `json:"data"`
	IsSuccess bool     `json:"is_success"`
	Messages  []string `json:"messages,omitempty"`
}

func NewSuccess(data any) Response {
	return Response{
		Data:      data,
		IsSuccess: true,
	}
}
func NewFailed(message string) Response {
	return Response{
		IsSuccess: false,
		Messages:  []string{message},
	}
}
