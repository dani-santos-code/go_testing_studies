package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func (t *testing.T) {
		d := Dictionary{"test":"this is just a test"}
		got, _ := d.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func (t *testing.T) {
		d := Dictionary{"test":"this is just a test"}
		_, err := d.Search("unknown")
		want := ErrNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, key given:%q", got, want, "test")
	}
}


func TestAdd(t *testing.T) {
	t.Run("new word", func (t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err:= d.Add(word, definition)

		assertError(t, err, nil)

		assertDefinition(t, d, word, definition)
	})

	t.Run("existing word", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
