package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	xpected := 4

	if sum != xpected {
		t.Errorf("expected '%d' but got '%d'",xpected, sum)
	}
}
// Example test case of 1 + 5
func ExampleAdd(){
	sum:= Add(1,5)
	fmt.Println(sum)
	// Output: 6
}