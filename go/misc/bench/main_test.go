package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkFmtSprintf(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = fmt.Sprintf("%s%s", str, "a")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer
	for n := 0; n < b.N; n++ {
		buf.WriteString("a")
	}
}

func BenchmarkStringsBuffer(b *testing.B) {
	var builder strings.Builder
	for n := 0; n < b.N; n++ {
		builder.WriteString("a")
	}
}
