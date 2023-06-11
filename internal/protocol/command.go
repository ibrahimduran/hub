package protocol

type ModeCommand struct {
}

func (c *ModeCommand) Metadata() Metadata {
	return Metadata{
		Name:        "mode",
		Description: "Switch to interactive mode",
	}
}

func (c *ModeCommand) Run(proto *Protocol) error {
	proto.Interactive = true
	return nil
}
