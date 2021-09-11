package godecoder

import "io"

type Decoder interface {
	Decode(r io.Reader, val interface{}) error
	Encode(w io.Writer, value interface{}) error
}
