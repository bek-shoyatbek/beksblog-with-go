package interfaces

type SlugReader interface {
	Read(slug string) (string, error)
}
