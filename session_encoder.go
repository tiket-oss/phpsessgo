package phpsessgo

import "github.com/imantung/phpsessgo/phpencode"

type SessionEncoder interface {
	Encode(session phpencode.PhpSession) (string, error)
	Decode(raw string) (phpencode.PhpSession, error)
}

type sessionEncoder struct {
	SessionEncoder
}

func (e *sessionEncoder) Encode(session phpencode.PhpSession) (string, error) {
	encoder := phpencode.NewPhpEncoder(session)
	return encoder.Encode()

}

func (e *sessionEncoder) Decode(raw string) (phpencode.PhpSession, error) {
	decoder := phpencode.NewPhpDecoder(raw)
	return decoder.Decode()
}
