package external

// Source ...
type Source struct {
	AdditionalName string
}

// Name ...
func (source Source) Name() string {
	return "external:" + source.AdditionalName
}
