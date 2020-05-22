package dictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	// ErrNotFound constant
	ErrNotFound = errors.New("could not find the word you were looking for")
	// ErrWordExists constant
	ErrWordExists = errors.New("this word has already been defined")
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
