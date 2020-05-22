package dictionary

// Dictionary type
type Dictionary map[string]string

// DicErr type
type DicErr string

func (e DicErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound constant
	ErrNotFound = DicErr("could not find the word you were looking for")
	// ErrWordExists constant
	ErrWordExists = DicErr("this word has already been defined")
	// ErrWordDoesNotExist constante
	ErrWordDoesNotExist = DicErr("cannot update the word because it does not exist")
)

// Search returns a term
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add allows for addition of words
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update allos for update
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
