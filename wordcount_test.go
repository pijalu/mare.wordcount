package main

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestSplitWord(t *testing.T) {
	input := "this is a text"
	expected := []string{"this", "is", "a", "text"}

	actual := splitWord(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestWordChannel(t *testing.T) {
	input := "this is a text"
	expected := splitWord(input)

	actualChan := wordChannel(strings.NewReader(input))
	for _, expect := range expected {
		actual := <-actualChan
		if expect != actual.(string) {
			t.Fatalf("Expected %v but got %v", expect, actual)
		}
	}
}

func benchmarkWordCount(wrk int, b *testing.B) {
	data, err := ioutil.ReadFile("corpus/t8.shakespeare.txt")
	if err != nil {
		b.Fatalf("Failed to read file: %v", err)
	}
	for n := 0; n < b.N; n++ {
		r := strings.NewReader(string(data))
		wordCount(r, wrk)
	}
}

func BenchmarkWordCount_w1(b *testing.B) { benchmarkWordCount(1, b) }
func BenchmarkWordCount_w2(b *testing.B) { benchmarkWordCount(2, b) }
func BenchmarkWordCount_w3(b *testing.B) { benchmarkWordCount(3, b) }
func BenchmarkWordCount_w4(b *testing.B) { benchmarkWordCount(4, b) }
