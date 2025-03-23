package maps

import "testing"

const (
	word       string = "test"
	definition string = "this is test"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, got, errWordNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, definition)
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("new word", func(t *testing.T) {
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Add(word, definition)

		assertError(t, err, errWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{}
	updDefinition := "this is test updated"
	dictionary.Add(word, definition)

	t.Run("update existing word", func(t *testing.T) {
		err := dictionary.Update("test", updDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updDefinition)
	})

	t.Run("update non existing word", func(t *testing.T) {
		err := dictionary.Update("test2", updDefinition)
		want := errWordNotFound
		assertError(t, err, want)
		assertDefinition(t, dictionary, "test2", updDefinition)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add(word, definition)

	t.Run("delete existing word", func(t *testing.T) {
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, errWordNotFound)
	})
}
