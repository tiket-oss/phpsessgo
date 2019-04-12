package phpsessgo

import "github.com/imantung/phpsessgo/phpencode"

type SessionEncoder interface {
	Encode(session phpencode.PhpSession) (string, error)
	Decode(raw string) (phpencode.PhpSession, error)
}
