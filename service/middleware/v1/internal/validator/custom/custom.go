package custom

type Validators interface {
	Empty(m map[string]string) bool
	TimeRange(m map[string]string) error
	PageRange(m map[string]string) error
}
