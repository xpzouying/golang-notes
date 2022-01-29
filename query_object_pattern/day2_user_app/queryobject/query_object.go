package queryobject

type Operator string

const (
	OpEqual   = "="
	OpGreater = ">"

	// more
)

type Field string

type Option func(v interface{}) interface{}
