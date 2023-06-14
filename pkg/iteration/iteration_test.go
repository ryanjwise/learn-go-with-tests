package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	character := "a-"
	times := 3
	response := Repeat(character, times)
	fmt.Println(response)
	// prints 'a-a-a-'
}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
