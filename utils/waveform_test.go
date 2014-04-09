package utils

import (
	"strings"
	"testing"
)

func TestGenerateSamples(t *testing.T) {

	var sample string = "0.00000,0.00000,0.00000,0.00000,0.00000,0.13814,0.14124,0.14326,0.14713,0.15060,0.15402,0.15601,0.15739,0.15901,0.16063,0.16206,0.16249,0.16379,0.16547,0.16731,0.16724,0.16891,0.16888,0.16603,0.16489,0.16460,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000,0.00000"
	var result []string = GenerateSamplesAsString("test.mp3", 5)
	if strings.Join(result, ",") != sample {
		t.Error("Parsing of file failed", result)
	}
}

func BenchmarkGenerateSamples(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < 10; n++ {
		GenerateSamplesAsFloat("test.mp3")
	}
}
