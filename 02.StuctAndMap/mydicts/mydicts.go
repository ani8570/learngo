package mydicts

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errAlreadyHas = errors.New("Already has it")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		d[key] = value
		return nil
	}
	return errAlreadyHas
}

// Update a word new word
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err == nil {
		d[key] = value
		return nil
	}
	return errCantUpdate
}

// Delete a word
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
