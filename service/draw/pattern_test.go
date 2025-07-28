package draw

import (
	"fmt"
	"testing"
)

func TestGeneratePatternFromText(t *testing.T) {
	text := GeneratePatternFromText("abcdefghijklmnopqrstuvwxyz")
	for _, line := range text {
		fmt.Println(line)
	}
}
