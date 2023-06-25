package pagination

type Comparison int

const (
	Equal Comparison = iota
	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
	NotEqual
	Contains
	StartsWith
	EndsWith
)

const (
	defaultSize = 10
)
