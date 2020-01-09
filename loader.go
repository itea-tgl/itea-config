package config

import (
	"path"
	"strings"
)

type ILoader interface {
	load(IProcessor, IExtractor) error
}

type Loader func() string

func (l Loader) load(p IProcessor, e IExtractor) error {
	file := l()
	data, err := p.Load(file)
	if err != nil {
		return err
	}
	e.Load(data, filename(file))
	return nil
}

func filename(file string) string {
	filename := path.Base(file)
	suffix := path.Ext(filename)
	return strings.TrimSuffix(filename, suffix)
}