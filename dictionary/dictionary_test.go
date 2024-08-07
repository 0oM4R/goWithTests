package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got,_ := dictionary.Search( "test")
		want := dictionary["test"]
		assertString(t, got, want)
	})
	t.Run("unknown word",func(t *testing.T)  {
		_,err := dictionary.Search("test")
		want := "could not find the word you were looking for"
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError
