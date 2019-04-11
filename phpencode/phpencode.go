package phpencode

import "github.com/imantung/phpsessgo/phptype"

const SEPARATOR_VALUE_NAME rune = '|'

type PhpSession map[string]phptype.PhpValue

// Decode decode string to PHP Session
func Decode(raw string) (PhpSession, error) {
	decoder := NewPhpDecoder(raw)
	return decoder.Decode()
}
