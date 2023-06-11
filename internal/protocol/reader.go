package protocol

import "context"

func ReadCmd(ctx context.Context, buffer chan byte, maxLength int) (string, error) {
	var cmd = make([]byte, maxLength)
	var char byte
	var head int

	done := ctx.Done()

	for {
		select {
		case <-done:
			return "", ctx.Err()
		case char = <-buffer:
		}

		if char == '\n' || char == ' ' {
			break
		}

		cmd[head] = char
		head++

		if head == CommandBufferSize {
			return "", ErrCommandBufferExceeded
		}
	}

	return string(cmd[:head]), nil
}
