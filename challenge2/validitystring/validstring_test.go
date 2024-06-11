package validitystring_test

import (
	"bytes"
	"decodetonum/validitystring"
	"log"
	"strings"
	"testing"
)

func TestCheckString(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"LL===", ""}, //อันแรกกรณีไม่มี error
		{"LLR=RRL=", "Error: Please ensure to enter 5 characters"},
		{"LLRMRRL=", "Error: Please ensure to enter 5 characters"},
		{"LLR=RRLZ", "Error: Please ensure to enter 5 characters"},
		{"LLAL=", "Error: String contains invalid character: A"},
		{"LL=RH", "Error: String contains invalid character: H"},
	}

	for _, test := range tests {
		var logBuf bytes.Buffer
		log.SetOutput(&logBuf)
		log.SetFlags(0)

		validitystring.CheckString(test.input)

		if !strings.Contains(logBuf.String(), test.expectedOutput) {
			t.Errorf("For input '%s', expected log output to contain '%s', but got '%s'", test.input, test.expectedOutput, logBuf.String())
		}
	}

	log.SetOutput(nil)
}
