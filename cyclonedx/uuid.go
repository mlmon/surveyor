package cyclonedx

import (
	"crypto/rand"
	"fmt"
)

var Random = rand.Read

func Uuid() (string, error) {
	b := make([]byte, 16)
	_, err := Random(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), err
}
