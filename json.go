package config

import (
	"encoding/json"
	"github.com/json-iterator/go"
)

var j jsoniter.API

func init() {
	j = jsoniter.Config{
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
	err = j.Unmarshal(dat, &d)
	if err != nil {
		return nil, err
	}
	jsonTrans(d)
	return d, nil
}

func jsonTrans(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		if v1, ok := v.(map[string]interface{}); ok {
			m[k] = jsonTrans(v1)
			continue
		}
		if v1, ok := v.(json.Number); ok {
			i, _ := v1.Int64()
			m[k] = int(i)
		}
	}
	return m
}

// JsonProcessor returns a Processor which load json config
func JsonProcessor() IProcessor {
	return &Json{}
}