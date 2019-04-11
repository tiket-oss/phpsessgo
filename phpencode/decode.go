package phpencode

// Decode decode string to PHP Session
func Decode(raw string) (PhpSession, error) {
	decoder := NewPhpDecoder(raw)
	return decoder.Decode()
}
