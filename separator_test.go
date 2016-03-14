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

	{"白玉粉", []string{"白", "玉", "粉"}},

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
	list string
	src  string
	dst  []string
}{
	{
		list: "",
		src:  "abc",
		dst:  []string{"a", "b", "c"},
	},

	{
		list: "",
		src:  "",
		dst:  []string{},
	},
	{
		list: ":.",
		src:  "",
		dst:  []string{},
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

	{
		list: "、",
		src:  "上、下、左、右",
		dst:  []string{"上", "下", "左", "右"},
	},
	{
		list: "］〕",
		src:  "あああああ］あああああああ〕あああああ",
		dst:  []string{"あああああ", "あああああああ", "あああああ"},
	},
}

func TestSeparate(t *testing.T) {
	for _, test := range separeteTests {
		s := NewSeparator(test.list)
		expect := test.dst
		actual := s.Separate(test.src)
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("%q.Separate(%q) = %q, want %q",
				s, test.src, actual, expect)
		}
	}
}

func TestSeparateTwice(t *testing.T) {
	list := ",."
	s := NewSeparator(list)

	var src string
	var expect, actual []string

	src = "aaa,bbb"
	expect = []string{"aaa", "bbb"}
	actual = s.Separate(src)
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("[1] %q.Separate(%q) = %q, want %q",
			s, src, actual, expect)
	}

	src = "ccc,ddd"
	expect = []string{"ccc", "ddd"}
	actual = s.Separate(src)
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("[2] %q.Separate(%q) = %q, want %q",
			s, src, actual, expect)
	}
}
