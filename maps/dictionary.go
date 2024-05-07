package main


type Dictionary map[string]string

type DictionaryError string

const (
	ErrorNotFound = DictionaryError("couldn't find the word you were looking for")
	ErrWordExists = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("cannot update word because it doesnt exist")
)

func(e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}


func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition

	case nil:
		return ErrWordExists

	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExist

	case nil:
		d[word]= definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func main() {

}
