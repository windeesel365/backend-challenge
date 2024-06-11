package wordcount

import (
	"reflect"
	"testing"
)

func TestCountMeats(t *testing.T) {
	text := "Beef beef chicken pork Beef"
	expected := map[string]int{"beef": 3, "chicken": 1, "pork": 1}

	result := CountMeats(text)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
