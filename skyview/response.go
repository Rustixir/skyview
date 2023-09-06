package skyview

type Response struct {
	BroadcastData interface{}
	Err           error
}

func RaiseError(err error) Response {
	return Response{
		Err: err,
	}
}

func Broadcast(data interface{}) Response {
	return Response{
		BroadcastData: data,
		Err:           nil,
	}
}

func Ok() Response {
	return Response{}
}
