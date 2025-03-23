package maps

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	errWordNotFound = DictionaryErr("could not find the word you were looking for")
	errWordExists   = DictionaryErr("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case errWordNotFound:
		d[key] = value
	case nil:
		return errWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key string, updVal string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = updVal
	case errWordNotFound:
		return err
	default:
		return err

	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
