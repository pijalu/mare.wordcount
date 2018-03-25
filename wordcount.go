package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pijalu/mare"
)

func splitWord(line string) []string {
	result := []string{}

	b := new(strings.Builder)
	for _, r := range []rune(line) {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			b.WriteRune(rune(r))
		} else if b.Len() > 0 {
			result = append(result, b.String())
			b.Reset()
		}
	}
	if b.Len() > 0 {
		result = append(result, b.String())
	}

	return result
}

func wordChannel(input io.Reader) chan interface{} {
	wordChan := make(chan interface{})
	// Read file per line and split word
	go func() {
		defer close(wordChan)
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			for _, word := range splitWord(scanner.Text()) {
				wordChan <- strings.ToLower(word)
			}
		}
	}()
	return wordChan
}

func wordCount(input io.Reader, workerCnt int) map[interface{}]interface{} {
	return mare.MaRe().MapWorker(workerCnt).InChannel(wordChannel(os.Stdin)).Map(func(input interface{}) []mare.MapOutput {
		return []mare.MapOutput{{Key: input, Value: 1}}
	}).Reduce(func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
}

func main() {
	for k, v := range wordCount(os.Stdin, 2) {
		fmt.Printf("%05d\t%s\n", v, k)
	}
}
