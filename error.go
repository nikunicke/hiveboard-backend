package hiveboard

const (
	ErrAuth = Error("Authorization token missing")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
