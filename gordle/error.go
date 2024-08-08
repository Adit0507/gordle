package gordle

// defines a sentinel error, sentinel error behave like constants
type corpusError string

func (c corpusError) Error() string {
	return string(c)
}
