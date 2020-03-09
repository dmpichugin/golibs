package errors

type String string

func (e String) Error() string {
	return string(e)
}
