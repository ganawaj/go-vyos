package vyos

import "errors"

var (

	ErrMethodNotSupported = errors.New("method not supported")
	ErrContextNil = errors.New("context must be non-nil")
	ErrInterfaceNil = errors.New("can not unmarshal into nil interface")
	ErrEmptyPath = errors.New("path cannot be empty")

)
