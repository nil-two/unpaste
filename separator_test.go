package main

import (
	"reflect"
	"testing"
)

var toDelimitersTests = []struct {
	list       string
	delimiters []string
}{
	{"abc", []string{"a", "b", "c"}},
	{"\n\t ", []string{"\n", "\t", " "}},
	{"<<--->>", []string{"<", "<", "-", "-", "-", ">", ">"}},
}

func TestToDelimiters(t *testing.T) {
	for _, test := range toDelimitersTests {
		expect := test.delimiters
		actual := toDelimiters(test.list)
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("toDelimiters(%q) = %q, want %q",
				test.list, actual, expect)
		}
	}
}
