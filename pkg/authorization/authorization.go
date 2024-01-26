package authorization

type Authorization interface {
	Validate(authorization string) error
	Fields(userAccess string) []string
}
