package auth

import (
	"github.com/ibrahimduran/hub/internal/protocol"
)

type RegisterCommand struct {
}

func (c *RegisterCommand) Metadata() protocol.Metadata {
	return protocol.Metadata{
		Name:        "register",
		Description: "Register user",
	}
}

func (c *RegisterCommand) Run(proto *protocol.Protocol) error {
	panic("Not implemented")
}
