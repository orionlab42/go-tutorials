package sayings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet(`James`)
	if s != "Welcome my dear James" {
		t.Error(`Expected:`, `"Welcome my dear James"`, `got`, s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}

// Commands

// godoc -http:=8080
// go test
// go test -bench .
// go test -cover
