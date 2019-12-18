package config

import "gopkg.in/yaml.v2"

type Yaml struct {
	Base
}

// load a yaml config file to map data
func (y Yaml) Load(file string) (map[string]interface{}, error) {
	dat, err := y.load(file)
	if err != nil {
		return nil, err
	}
	var d map[string]interface{}
	err = yaml.Unmarshal(dat, &d)
	if err != nil {
		return nil, err
	}

	return yamlTrans(d), nil
}

//transform the interface{} of map[interface{}] interface{} to map[string]interface{}
func yamlTrans(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		if v1, ok := v.(map[interface{}]interface{}); ok {
			mm := make(map[string]interface{})
			for k1, v2 := range v1 {
				mm[k1.(string)] = v2
			}
			m[k] = yamlTrans(mm)
		}
	}
	return m
}

// YamlProcessor returns a Processor which load yaml config
func YamlProcessor() IProcessor {
	return &Yaml{}
}