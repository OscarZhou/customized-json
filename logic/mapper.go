package logic

// Mapper is a
type Mapper interface {
	ToKeys() ([]string, error)
}
