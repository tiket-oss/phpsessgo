package phpsessgo

import "github.com/tiket-oss/phpsessgo/phpencode"

type SessionEncoder interface {
	Encode(session phpencode.PhpSession) (string, error)
	Decode(raw string) (phpencode.PhpSession, error)
}
