package phpencode

import (
	"bytes"
	"fmt"

	"github.com/imantung/phpsessgo/phpserialize"
)

type PhpEncoder struct {
	data    PhpSession
	encoder *phpserialize.Serializer
}

func NewPhpEncoder(data PhpSession) *PhpEncoder {
	return &PhpEncoder{
		data:    data,
		encoder: phpserialize.NewSerializer(),
	}
}

func (self *PhpEncoder) SetEncodeFunc(f phpserialize.EncodeFunc) {
	self.encoder.SetEncodeFunc(f)
}

func (self *PhpEncoder) Encode() (string, error) {
	if self.data == nil {
		return "", nil
	}
	var (
		err error
		val string
	)
	buf := bytes.NewBuffer([]byte{})

	for k, v := range self.data {
		buf.WriteString(k)
		buf.WriteRune(SEPARATOR_VALUE_NAME)
		if val, err = self.encoder.Encode(v); err != nil {
			err = fmt.Errorf("php_session: error during encode value for %q: %v", k, err)
			break
		}
		buf.WriteString(val)
	}

	return buf.String(), err
}
