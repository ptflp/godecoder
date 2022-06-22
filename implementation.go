package godecoder

import (
	jsoniter "github.com/json-iterator/go"
	"io"
)

var defaultConfig = jsoniter.ConfigCompatibleWithStandardLibrary

var decoderSingleton *Decode

type Decode struct {
	api jsoniter.API
}

func NewDecoder(args ...jsoniter.Config) Decoder {
	conf := defaultConfig
	if len(args) == 0 && decoderSingleton == nil {
		decoderSingleton = &Decode{
			api: conf,
		}
		return decoderSingleton
	}
	if len(args) > 0 {
		conf = args[0].Froze()
	}

	return &Decode{
		api: conf,
	}
}

func (d *Decode) Decode(r io.Reader, val interface{}) error {
	var decoder = d.api.NewDecoder(r)
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

func (d *Decode) Encode(w io.Writer, value interface{}) error {
	return d.api.NewEncoder(w).Encode(value)
}
