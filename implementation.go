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
	if len(args) < 1 && decoderSingleton == nil {
		decoderSingleton = &Decode{
			api: defaultConfig,
		}
		return decoderSingleton
	}

	return &Decode{
		api: args[0].Froze(),
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
