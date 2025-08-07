package main

import (
	"log"
	"strings"
	"testing"
)

func testName(t *testing.T) {
	name1 := reverseName(strings.TrimSpace("William"))
	expected1 := "mailliW"

	if !strings.EqualFold(name1, expected1) {
		t.Errorf("Response from reverseName is unexpected value. Got [%s], expected [%s]", name1, expected1)
	}

	name2 := reverseName(strings.TrimSpace("Mister"))
	expected2 := "retsiM"
	log.Println(name2, expected2, len(name2), len(expected2))
	if !strings.EqualFold(name2, expected2) {
		t.Errorf("Response from reverseName is unexpected value. Got [%s], expected [%s]", name2, expected2)
	}
}
