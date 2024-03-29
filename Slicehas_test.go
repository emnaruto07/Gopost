package main

import "testing"

func TestCommandHas(t *testing.T) {
	hay := []string{"a", "b", "c"}
	expected := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"e": false,
	}

	for needle, expectation := range expected {
		output := CommandHas(hay, needle)
		if output != expectation {
			t.Fatalf("for input=%s, expected=%v, but received=%v", needle, expectation, output)
		}
	}
}
