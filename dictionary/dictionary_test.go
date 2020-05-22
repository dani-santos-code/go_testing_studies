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
	d := Dictionary{}
	d.Add("car", "A means of transportation")

	word := "car"
	definition := "A means of transportation"

	assertDefinition(t, d, word, definition)
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
