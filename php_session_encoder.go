package phpsessgo

import "github.com/imantung/phpsessgo/phpencode"

type PHPSessionEncoder struct {
	SessionEncoder
}

func (e *PHPSessionEncoder) Encode(session phpencode.PhpSession) (string, error) {
	encoder := phpencode.NewPhpEncoder(session)
	return encoder.Encode()

}

func (e *PHPSessionEncoder) Decode(raw string) (phpencode.PhpSession, error) {
	decoder := phpencode.NewPhpDecoder(raw)
	return decoder.Decode()
}
