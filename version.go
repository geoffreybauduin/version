package version

type Version interface {
	Is(string) bool
	Number() interface{}
}

func New(versionString string) (Version, error) {
	v := &versionInt{}
	return v, v.parse(versionString)
}
