package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	expect := 6
	if sum != expect {
		t.Errorf("expected '%d' got '%d", expect, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//	Output: 6

}
