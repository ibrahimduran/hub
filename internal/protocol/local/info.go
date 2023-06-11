package local

import (
	"os"

	"github.com/ibrahimduran/hub/internal/protocol"
)

type Info struct {
}

func (c *Info) Metadata() protocol.Metadata {
	return protocol.Metadata{
		Name:        "info",
		Description: "Display info",
	}
}

func (c *Info) Run(protocol *protocol.Protocol) error {
	hostname, err := os.Hostname()

	if err != nil {
		return err
	}

	protocol.Write([]byte(hostname + "\n"))

	return nil
}
