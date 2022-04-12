package handle

import (
	"context"
	"io"
)

type Handel interface {
	onRead(c io.Reader, ctx context.Context)
	//OnWrite(w io.Writer, ctx context.Context)
}
