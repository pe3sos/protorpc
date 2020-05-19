package protorpc

// StatusCode numeric representation of the request handling result conditions
type StatusCode int32

var BasicError = struct {
	BadCallPath    StatusCode
	InvalidPayload StatusCode
}{
	BadCallPath:    1,
	InvalidPayload: 2,
}

// StatusError provides not only description but also int status code of the error.
type StatusError interface {
	error
	WithID(string) StatusError
	ID() string
	StatusCode() StatusCode
}

type RPCError struct {
	ErrID string
	Msg   string
	Code  StatusCode
}

func (e RPCError) WithID(id string) StatusError {
	e.ErrID = id
	return e
}

func (e RPCError) ID() string {
	return e.ErrID
}

func (e RPCError) Error() string {
	return e.Msg
}

func (e RPCError) StatusCode() StatusCode {
	return e.Code
}
