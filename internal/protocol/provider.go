package protocol

type Provider interface {
	Execute(cmd string, protocol *Protocol) error
	CanExecute(cmd string, protocol *Protocol) bool
	List() []Metadata
}

type Metadata struct {
	Name        string
	Description string
}
