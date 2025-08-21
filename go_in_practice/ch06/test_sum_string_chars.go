package main

import (
	"testing"
)

func TestSummedRuneCodes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int16
	}{
		{"test 1", "i am trying things", 1729},
		{"test 2", "doing my best to find a way to break this!", 3772},
		{"test 3", "even adding emojis or unicode doesn't break it, I'm sure I'm fine figure", 6065},
	}

	t.Run("my rune sum tests", func(t *testing.T) {
		if got := summedRuneCodes(test.input); got != test.expected {
			t.Errorf("expected %d, got %d", test.expected, got)
		}
	})
}
