package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound

	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

var (
	ErrWordNotFound = DictionaryErr("could not find what you are looking for")
	ErrWordExists   = DictionaryErr("word already exists")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
