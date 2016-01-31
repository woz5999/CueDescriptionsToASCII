package elements

// CueElement ... interface for all CueElements
type CueElement interface {
	Convert() string
	Validate() bool
	SetValue()
}
