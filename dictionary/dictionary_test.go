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
	d.Add("name", "Dani")

	want := "Dani"
	got, err := d.Search("name")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
