package main

// a Map is a pointer to a runtime.hmap structure. So when passed by value the pointer is copied but still points to the same underlying hmap
type Dictionary map[string]string

// Search returns the definion of a word if it is present, or an error (ErrNotFound) if the word hasn't been defined
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add sets the word and definition in the dictionary.
//If the word is already defined ErrWordDoesNotExist is returned
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordDoesNotExist
	default:
		return err

	}
	return nil
}
//Update replaces the definition of a word if it exists in the dictionary.
//If the word is not defined ErrWordDoesNotExist is returned
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err

	}
	return nil
}

//Delete removes a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}
