package models

import (
	"testing"
)

func TestFromByteSliceOK(t *testing.T) {
	testData := []byte("abcd")
	hist := AsciiHist{}
	hist.FromByteSlice(testData)
	if hist != [128]int{97: 1, 98: 1, 99: 1, 100: 1} {
		t.Error("TestFromByteSliceOK failed")
	}
}

func TestFromByteSliceEmpty(t *testing.T) {
	testData := []byte("")
	hist := AsciiHist{}
	hist.FromByteSlice(testData)
	if hist != [128]int{} {
		t.Error("TestFromByteSliceEmpty failed")
	}
}

func TestFromByteSliceMultiSymbol(t *testing.T) {
	testData := []byte("aaabb")
	hist := AsciiHist{}
	hist.FromByteSlice(testData)
	if hist != [128]int{97: 3, 98: 2} {
		t.Error("TestFromByteSliceMultiSymbol failed")
	}
}

func TestAddTwoHist(t *testing.T) {
	firstTestData := []byte("aaabb")
	secondTestData := []byte("bccc")
	firstHist := AsciiHist{}
	secondHist := AsciiHist{}
	firstHist.FromByteSlice(firstTestData)
	secondHist.FromByteSlice(secondTestData)
	firstHist.Add(secondHist)
	if firstHist != [128]int{97: 3, 98: 3, 99: 3} {
		t.Error("TestAddTwoEmptyHist failed")
	}
}

func TestAddTwoEmptyHist(t *testing.T) {
	firstHist := AsciiHist{}
	secondHist := AsciiHist{}
	firstHist.Add(secondHist)
	if firstHist != [128]int{} {
		t.Error("TestAddTwoEmptyHist failed")
	}
}

func TestAddEmptyAndNotEmptyHist(t *testing.T) {
	firstTestData := []byte("aaabb")
	firstHist := AsciiHist{}
	secondHist := AsciiHist{}
	firstHist.FromByteSlice(firstTestData)
	firstHist.Add(secondHist)
	if firstHist != [128]int{97: 3, 98: 2} {
		t.Error("TestAddTwoEmptyHist failed")
	}
}
