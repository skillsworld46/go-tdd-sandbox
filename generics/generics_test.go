package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {

		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on string", func(t *testing.T) {
		AssertEqual(t, "string", "string")
		AssertNotEqual(t, "string", "string no. 2")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %d", got)
	}
}
func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		intStack := new(Stack[int])

		AssertTrue(t, intStack.IsEmpty())

		intStack.Push(123)
		AssertFalse(t, intStack.IsEmpty())

		intStack.Push(456)
		val, _ := intStack.Pop()
		AssertEqual(t, val, 456)
		val, _ = intStack.Pop()
		AssertEqual(t, val, 123)

		AssertTrue(t, intStack.IsEmpty())
		intStack.Push(1)
		intStack.Push(2)
		firstNum, _ := intStack.Pop()
		secondNum, _ := intStack.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
