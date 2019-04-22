package keycloak

import "fmt"

// Error represents an API error
type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %d", e.Message, e.Code)
}
