package protocol

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	CommandBufferSize = 100
)

type Protocol struct {
	Providers         []Provider
	Stream            io.ReadWriteCloser
	BufferBacklogSize int
	BufferSize        int
	Interactive       bool

	ctx    context.Context
	buffer chan byte
}

func (p *Protocol) Run(ctx context.Context) {
	p.buffer = make(chan byte, p.BufferBacklogSize)

	var cancel context.CancelCauseFunc
	p.ctx, cancel = context.WithCancelCause(ctx)

	go func() {
		buf := make([]byte, p.BufferSize)

		for {
			readBytes, err := p.Stream.Read(buf)

			if errors.Is(err, io.EOF) {
				cancel(ErrSocketClosed)
				break
			}

			if err != nil {
				panic(err)
			}

			for i := 0; i < readBytes; i++ {
				p.buffer <- buf[i]
			}
		}
	}()

	p.onInit()

	for p.ctx.Err() == nil {
		p.onPrompt()
	}

	p.onClean()
}

func (p *Protocol) Write(data []byte) (int, error) {
	return p.Stream.Write(data)
}

func (p *Protocol) Error(err error) {
	p.Stream.Write([]byte(err.Error() + "\n"))
}

func (p *Protocol) Close() {

	p.Stream.Close()
}

func (p *Protocol) onInit() {
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "unknown"
	}

	p.Write([]byte("dizin hub @ " + hostname + "\n"))

	fmt.Println("client connected")
}

func (p *Protocol) onPrompt() {
	if p.Interactive {
		p.Write([]byte("> "))
	}

	cmd, err := ReadCmd(p.ctx, p.buffer, CommandBufferSize)

	if len(cmd) == 0 {
		return
	}

	if err != nil {
		if !errors.Is(err, context.Canceled) {
			p.Error(err)
		}

		return
	}

	for _, provider := range p.Providers {
		err := provider.Execute(cmd, p)

		if errors.Is(err, ErrCanNotExecute) {
			continue
		}

		if err != nil {
			p.Error(err)
		}

		return
	}

	p.Error(ErrCommandNotFound)
}

func (p *Protocol) onClean() {
	fmt.Println(fmt.Errorf("client disconnected: %w", context.Cause(p.ctx)))
}
