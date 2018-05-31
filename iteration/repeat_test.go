package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeatedOutput := repeat("a", 5)
	fmt.Println(repeatedOutput)
	//Output: aaaaa
}

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := repeat("a", 5)

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 10)
	}
}
