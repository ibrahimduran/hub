package local

import (
	"fmt"
	"text/tabwriter"

	"github.com/ibrahimduran/hub/internal/protocol"
)

type Help struct {
}

func (c *Help) Metadata() protocol.Metadata {
	return protocol.Metadata{
		Name:        "help",
		Description: "Display help",
	}
}

func (c *Help) Run(proto *protocol.Protocol) error {
	count := 0
	table := tabwriter.NewWriter(proto, 1, 1, 1, ' ', 0)

	for _, provider := range proto.Providers {
		for _, cmd := range provider.List() {
			table.Write([]byte(cmd.Name + "\t" + cmd.Description + "\n"))
			count++
		}
	}

	table.Flush()
	proto.Write([]byte(fmt.Sprintf("Found %d commands.\n", count)))

	return nil
}
