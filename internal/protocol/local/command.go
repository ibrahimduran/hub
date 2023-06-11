package local

import "github.com/ibrahimduran/hub/internal/protocol"

type CommandCallback func(proto *protocol.Protocol) error

type Command interface {
	Metadata() protocol.Metadata
	Run(protocol *protocol.Protocol) error
}
