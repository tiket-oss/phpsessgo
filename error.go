package phpsessgo

const (
	ErrNotSupported = Error("Not supported")
)

// Error is phpsessgo error model
type Error string

func (e Error) Error() string {
	return string(e)
}
