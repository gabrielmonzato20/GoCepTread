package entity

type Response interface {
}
type ResponseEntity struct {
	ApiResponse string
	Response    Response
}

func NewResponseEntity(api string, responseDto Response) *ResponseEntity {
	return &ResponseEntity{
		ApiResponse: api,
		Response:    responseDto,
	}
}
