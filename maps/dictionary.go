package maps

// message errors for Dictionary
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr represents the error messages on Dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary is a type that wraps a map of strings
type Dictionary map[string]string

// Search is a method of dictionary which returns a value given a key
func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add is method to add new words to a Dictionary
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

// Update is a method to update existing words in a Dictionary
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

// Delete is a method to remove existing words from a Dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
