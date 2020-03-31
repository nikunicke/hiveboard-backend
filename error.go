package hiveboard

const (
	ErrAuth       = Error("Authorization token missing")
	UserNotFound  = Error("User does not exist")
	EventNotFound = Error("Event does not exist")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
