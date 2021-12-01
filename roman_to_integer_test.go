package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var numerals = map[string]int {
		"III": 3,
		"IV": 4,
		"CIX": 109,
	}

	for key, value := range numerals {
		retVal := romanToInt(key)
        if retVal != value {
			t.Fatalf("Invalid answer. Numeral given: %v. Returned value %v. Expected returned value to be: %v.\n", key, retVal, value)
		}
	}
}