package validate

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return v[0].Message
}

type ValidationError struct {
	Tag     string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path"`
	Value   any    `json:"value"`
}

func (v *ValidationError) Error() string {
	return v.Message
}
