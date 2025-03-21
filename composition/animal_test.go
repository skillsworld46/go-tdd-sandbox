package composition

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWarmBlood(t *testing.T) {
	want := true
	d := Dog{Breed: "caigou", Animal: Animal{Name: "xiao hei"}}
	fmt.Println("Hello, I am saying", d.Name, "is a", reflect.TypeOf(d).Name())
	got := d.IsWarmBlood()
	if got != want {
		t.Errorf("Dog is not cold blood")
	}
}
