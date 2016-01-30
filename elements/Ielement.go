package elements

type CueElement interface {
	Convert() string
	Validate() bool
}
