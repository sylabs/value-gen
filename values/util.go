package values

import (
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomSecret(l int) string {
	var s strings.Builder
	s.Grow(l)

	for i := 0; i < l; i++ {
		r := rand.Intn(len(charset))
		s.WriteByte(charset[r])
	}
	return s.String()
}
