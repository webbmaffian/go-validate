package validate

type ValidationError struct {
	Tag     string
	Message string
	Path    string
	Value   any
}
