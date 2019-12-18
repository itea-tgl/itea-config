package config

import (
	"github.com/json-iterator/go"
)

var json jsoniter.API

func init() {
	json = jsoniter.Config{
		EscapeHTML:             false,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		UseNumber:				true,
	}.Froze()
}

type Json struct {
	Base
}

// load a json config file to map data
func (y Json) Load(file string) (map[string]interface{}, error) {
	dat, err := y.load(file)
	if err != nil {
		return nil, err
	}
	var d map[string]interface{}
	err = json.Unmarshal(dat, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// JsonProcessor returns a Processor which load json config
func JsonProcessor() IProcessor {
	return &Json{}
}