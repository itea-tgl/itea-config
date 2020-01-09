package config

import (
	"gopkg.in/ini.v1"
	"strings"
)

type Ini struct {

}

// load a ini config file to map data
func (y Ini) Load(file string) (map[string]interface{}, error) {
	dat, err := ini.Load(file)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	for _, section := range dat.Sections() {
		data[section.Name()] = sectionExtract(section)
	}
	return data, nil
}

// extract ini section config
func sectionExtract(section *ini.Section) map[string]interface{} {
	item := make(map[string]interface{})
	for _, k := range section.KeyStrings() {
		if k == "" {
			continue
		}
		kArr := strings.Split(k, ".")
		child := iniTrans(kArr, section.Key(k).String())
		if v1, ok := item[kArr[0]]; ok {
			if v2, ok := v1.(map[string]interface{}); ok {
				if v3, ok := child.(map[string]interface{}); ok {
					item[kArr[0]] = mapMerge(v2, v3)
					continue
				}
			}
		}
		item[kArr[0]] = child
	}
	return item
}

// transform ini section key
func iniTrans(k []string, v string) interface{} {
	if len(k) == 1 {
		return v
	}
	return map[string]interface{}{
		k[1]: iniTrans(k[1:], v),
	}
}

func mapMerge(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// IniProcessor returns a Processor which load ini config
func IniProcessor() IProcessor {
	return &Ini{}
}