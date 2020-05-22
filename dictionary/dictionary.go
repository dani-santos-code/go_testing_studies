package dictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

// ErrNotFound constant
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search returns a term
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add allows for addition of words
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
