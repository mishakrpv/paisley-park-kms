package configuration

var Config Configuration

type Configuration interface {
	Get(key string) string
}
