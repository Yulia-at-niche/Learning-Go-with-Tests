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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

var (
	ErrWordNotFound     = DictionaryErr("could not find the word you are looking for")
	ErrWordExists       = DictionaryErr("word already exists, cannot add")
	ErrWordDoesNotExist = DictionaryErr("word does not exist, cannot perform operation")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
