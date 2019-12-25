package config

type ILoader interface {
	load() (map[string]interface{}, error)
}

type Loader func() (map[string]interface{}, error)

func (l Loader) load() (map[string]interface{}, error) {
	return l()
}