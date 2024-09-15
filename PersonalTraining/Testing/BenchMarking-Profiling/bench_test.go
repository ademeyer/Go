package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var benchText = "Don't communicate by sharing memory, share memory by communicating."

func compareSlice(s1, s2 []string) bool {
	result := false
	if len(s1) != len(s2) {
		return result
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			result = false
			break
		}
	}
	return result
}

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(benchText)
		require.Equal(b, 10, len(tokens))
	}
}

// go test -v -bench . -run NONE
// go test -v -bench . -run NONE -cpuprofile=cpu.prof
// go tool pprof -http=:8080 cpu.prof
func TestTokenize(t *testing.T) {
	text := "who's on first?"
	expected := []string{"who", "s", "on", "first"}
	tokens := Tokenize(text)

	if !compareSlice(expected, tokens) {
		t.Fatalf("%v != %v", expected, tokens)
	}
}
