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
	date := make(map[string]interface{})
	for _, section := range dat.Sections() {
		date[section.Name()] = sectionExtract(section)
	}
	return date, nil
}

// extract ini section config
func sectionExtract(section *ini.Section) map[string]interface{} {
	item := make(map[string]interface{})
	for _, k := range section.KeyStrings() {
		if k == "" {
			continue
		}
		kArr := strings.Split(k, ".")
		item[kArr[0]] = iniTrans(kArr, section.Key(k).String())

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

// IniProcessor returns a Processor which load ini config
func IniProcessor() IProcessor {
	return &Ini{}
}