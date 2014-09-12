package gcss

// atRule represents an at-rule of CSS.
type atRule struct {
	elementBase
	sels []*selector
	decs []*declaration
}
