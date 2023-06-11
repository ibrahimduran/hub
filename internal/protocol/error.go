package protocol

import "errors"

var (
	ErrCommandNotFound       = errors.New("command not found")
	ErrCommandBufferExceeded = errors.New("command buffer length exceeded")
	ErrSocketClosed          = errors.New("socket closed")
	ErrCanNotExecute         = errors.New("can not execute this command")
)
