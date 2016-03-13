package main

import (
	"reflect"
	"testing"
)

var toDelimitersTests = []struct {
	list       string
	delimiters []string
}{
	{"", []string{""}},

	{"abc", []string{"a", "b", "c"}},
	{"\n\t ", []string{"\n", "\t", " "}},
	{"<<--->>", []string{"<", "<", "-", "-", "-", ">", ">"}},

	{"\\a\\b\\c", []string{"a", "b", "c"}},
	{"\\\\\\t\\n\\0", []string{"\\", "\t", "\n", ""}},
	{"\\", []string{""}},
	{"\\\\\\", []string{"\\"}},
	{"\\\\\\\\\\", []string{"\\", "\\"}},
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

var separeteTests = []struct {
	isSerial bool
	list     string
	src      string
	dst      []string
}{
	{
		list: "",
		src:  "abc",
		dst:  []string{"a", "b", "c"},
	},

	{
		list: "\t",
		src:  "aaa\tbbb\tccc",
		dst:  []string{"aaa", "bbb", "ccc"},
	},
	{
		list: ",",
		src:  "1,,2,,3",
		dst:  []string{"1", "", "2", "", "3"},
	},

	{
		list: ":.",
		src:  "aaa:bbb.ccc:ddd.eee",
		dst:  []string{"aaa", "bbb", "ccc", "ddd", "eee"},
	},
	{
		list: ">])}",
		src:  "a>bb]c)dd}e>ff]g",
		dst:  []string{"a", "bb", "c", "dd", "e", "ff", "g"},
	},
}

func TestSeparate(t *testing.T) {
	for _, test := range separeteTests {
		s := NewSeparator(test.isSerial, test.list)
		expect := test.dst
		actual := s.Separate(test.src)
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("%q.Separate(%q) = %q, want %q",
				s, test.src, actual, expect)
		}
	}
}
