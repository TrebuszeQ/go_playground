package main

import (
	"testing"
)

func FuzzSummedRuneCodes(f *testing.F) {
	tests := []string{"i am trying things", "doing my best to find a way to break this!", "even adding emojis or unicode doesn't break it, I'm sure I'm fine figure"}
	for t := range tests {
		f.Add(tests[t])
	}
	f.Fuzz(func(t *testing.T, seed string) {
		got := summedRuneCodes(seed)
		if got < 0 {
			t.Errorf("how did this happen? somehow we got %d from string %s", got, seed)
		}
	})
}	