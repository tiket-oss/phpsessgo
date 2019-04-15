package phpencode

import (
	"bytes"
	"io"
	"strings"

	"github.com/tiket-oss/phpsessgo/phpserialize"
	"github.com/tiket-oss/phpsessgo/phptype"
)

type PhpDecoder struct {
	source  *strings.Reader
	decoder *phpserialize.Unserializer
}

func NewPhpDecoder(phpSession string) *PhpDecoder {
	decoder := &PhpDecoder{
		source:  strings.NewReader(phpSession),
		decoder: phpserialize.NewUnserializer(""),
	}
	decoder.decoder.SetReader(decoder.source)
	return decoder
}

func (self *PhpDecoder) SetDecodeFunc(f phpserialize.DecodeFunc) {
	self.decoder.SetDecodeFunc(f)
}

func (self *PhpDecoder) Decode() (PhpSession, error) {
	var (
		name  string
		err   error
		value phptype.Value
	)
	res := make(PhpSession)

	for {
		if name, err = self.readName(); err != nil {
			break
		}
		if value, err = self.decoder.Decode(); err != nil {
			break
		}
		res[name] = value
	}

	if err == io.EOF {
		err = nil
	}
	return res, err
}

func (self *PhpDecoder) readName() (string, error) {
	var (
		token rune
		err   error
	)
	buf := bytes.NewBuffer([]byte{})
	for {
		if token, _, err = self.source.ReadRune(); err != nil || token == SEPARATOR_VALUE_NAME {
			break
		} else {
			buf.WriteRune(token)
		}
	}
	return buf.String(), err
}
