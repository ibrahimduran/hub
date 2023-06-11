package local

import (
	"strings"

	"github.com/ibrahimduran/hub/internal/protocol"
)

type LocalProvider struct {
	callbacks map[string]CommandCallback
	metadatas map[string]protocol.Metadata
}

func (p *LocalProvider) Register(cmd Command) {
	metadata := cmd.Metadata()
	name := normalizeCommandName(metadata.Name)

	_, exists := p.callbacks[name]

	if exists {
		panic("Command already exists")
	}

	if p.callbacks == nil {
		p.callbacks = make(map[string]CommandCallback)
	}

	if p.metadatas == nil {
		p.metadatas = make(map[string]protocol.Metadata)
	}

	p.callbacks[name] = cmd.Run
	p.metadatas[name] = metadata
}

func (p *LocalProvider) Add(cmd string, run CommandCallback) {

}

func (p *LocalProvider) Execute(cmd string, proto *protocol.Protocol) error {
	run, exists := p.callbacks[normalizeCommandName(cmd)]

	if !exists {
		return protocol.ErrCanNotExecute
	}

	return run(proto)
}

func (p *LocalProvider) CanExecute(cmd string, proto *protocol.Protocol) bool {
	_, exists := p.callbacks[normalizeCommandName(cmd)]
	return exists
}

func (p *LocalProvider) List() []protocol.Metadata {
	list := make([]protocol.Metadata, len(p.metadatas))
	i := 0
	for _, v := range p.metadatas {
		list[i] = v
		i++
	}
	return list
}

func normalizeCommandName(cmd string) string {
	return strings.ToUpper(cmd)
}
