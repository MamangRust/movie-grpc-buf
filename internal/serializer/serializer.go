package serializer

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

type JSON interface {
	JSON() ([]byte, error)
	FromJSON(data []byte) error
}

type YAML interface {
	YAML() ([]byte, error)
	FromYAML(data []byte) error
}

type API[T any] interface {
	API() T
	FromAPI(in T)
}
