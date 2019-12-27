package config

type ILoader interface {
	load(IExtractor) error
}

type Loader func(e IExtractor) error

func (l Loader) load(e IExtractor) error {
	return l(e)
}