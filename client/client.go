package client

import "fmt"

func wrapError(msg string, originalError error) error {
	return fmt.Errorf("%s: %v", msg, originalError)
}
