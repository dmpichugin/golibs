package uniq

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUUIDRegexp(t *testing.T) {
	re, err := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$")
	require.NoError(t, err, "compile regular")
	require.True(t, re.MatchString(UUID()))
}

func BenchmarkUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}
