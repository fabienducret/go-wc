package wc_test

import (
	"strings"
	"testing"
	"wc/wc"
)

func TestWC(t *testing.T) {
	t.Run("get wc information", func(t *testing.T) {
		// Arrange
		str := "words in line 1\nwords in line 2"
		reader := strings.NewReader(str)

		// Act
		counts := wc.CountsFor(reader)

		// Assert
		assert(t, counts.Lines, 2)
		assert(t, counts.Words, 8)
		assert(t, counts.Characters, 30)
	})
}

func assert(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("error got %d, want %d", got, want)
	}
}
