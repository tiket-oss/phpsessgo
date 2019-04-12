package phpsessgo

import "github.com/google/uuid"

// UUIDCreator generate session ID using UUID V4
type UUIDCreator struct {
	SessionIDCreator
}

func (c *UUIDCreator) CreateSID() string {
	return uuid.New().String()
}
