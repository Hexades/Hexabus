package bus

type RequestResponseEvent struct {
	responseChannel chan Response
}

func (e *RequestResponseEvent) getChannel() chan Response {
	if e.responseChannel == nil {
		e.responseChannel = make(chan Response, 1)
	}
	return e.responseChannel
}

func (e *RequestResponseEvent) Send(val Response) {
	e.getChannel() <- val
}

func (e *RequestResponseEvent) Receive() Response {
	return <-e.getChannel()
}
func NewResponse(value any, err error) Response {
	return Response{value, err}
}

type Response struct {
	Value any
	Err   error
}
