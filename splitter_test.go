package splitter_test

import (
	"github.com/ysugimoto/splitter"
	"testing"
)

func compareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func compareBytesSlice(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !compareBytesSliceChild(v, b[i]) {
			return false
		}
	}

	return true
}

func compareBytesSliceChild(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestSplitString(t *testing.T) {
	str := "This is sample of testcase"
	actual := splitter.SplitString(str, " ")
	expect := []string{"This", "is", "sample", "of", "testcase"}

	if !compareSlice(actual, expect) {
		t.Error("SplitString failed unexpected result")
	}
}

func TestSplitStringSepStringInSingleQuoteSequence(t *testing.T) {
	str := "This is sample of testcase | of seprapte '|' in sequence"
	actual := splitter.SplitString(str, "|")
	expect := []string{"This is sample of testcase ", " of seprapte '|' in sequence"}

	if !compareSlice(actual, expect) {
		t.Error("SplitString failed unexpected result when sep string in single quote sequence")
	}
}

func TestSplitStringSepStringInDoubleQuoteSequence(t *testing.T) {
	str := "This is sample of testcase | of seprapte \"|\" in sequence"
	actual := splitter.SplitString(str, "|")
	expect := []string{"This is sample of testcase ", " of seprapte \"|\" in sequence"}

	if !compareSlice(actual, expect) {
		t.Error("SplitString failed unexpected result when sep string in double quote sequence")
	}
}

func TestSplitStringSepStringInMixedSequence(t *testing.T) {
	str := "This is sample of testcase | of seprapte ''\"|' in sequence"
	actual := splitter.SplitString(str, "|")
	expect := []string{"This is sample of testcase ", " of seprapte ''\"", "' in sequence"}

	if !compareSlice(actual, expect) {
		t.Error("SplitString failed unexpected result when sep string in mixed quote sequence")
	}
}

func TestSplitBytes(t *testing.T) {
	str := "Lorem ipsum dolor sit amet"
	actual := splitter.SplitBytes([]byte(str), []byte(" "))
	expect := [][]byte{
		[]byte("Lorem"),
		[]byte("ipsum"),
		[]byte("dolor"),
		[]byte("sit"),
		[]byte("amet"),
	}

	if !compareBytesSlice(actual, expect) {
		t.Error("SplitBytes failed unexpected result")
	}
}

func TestSplitBytesSepByteInSingleQuoteSequence(t *testing.T) {
	str := "This is sample of testcase | of seprapte '|' in sequence"
	actual := splitter.SplitBytes([]byte(str), []byte("|"))
	expect := [][]byte{
		[]byte("This is sample of testcase "),
		[]byte(" of seprapte '|' in sequence"),
	}

	if !compareBytesSlice(actual, expect) {
		t.Error("SplitBytes failed unexpected result when sep bytes in single quote sequence")
	}
}

func TestSplitBytesSepByteInSingleQuoteSequenceDoubleQuoteSequence(t *testing.T) {
	str := "This is sample of testcase | of seprapte \"|\" in sequence"
	actual := splitter.SplitBytes([]byte(str), []byte("|"))
	expect := [][]byte{
		[]byte("This is sample of testcase "),
		[]byte(" of seprapte \"|\" in sequence"),
	}

	if !compareBytesSlice(actual, expect) {
		t.Error("SplitBytes failed unexpected result when sep bytes in double quote sequence")
	}
}

func TestSplitMixedSepBytesInMixedSequence(t *testing.T) {
	str := "This is sample of testcase | of seprapte ''\"|' in sequence"
	actual := splitter.SplitBytes([]byte(str), []byte("|"))
	expect := [][]byte{
		[]byte("This is sample of testcase "),
		[]byte(" of seprapte ''\""),
		[]byte("' in sequence"),
	}

	if !compareBytesSlice(actual, expect) {
		t.Error("SplitBytes failed unexpected result when sep string in mixed quote sequence")
	}
}
