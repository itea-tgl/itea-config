package config

import (
	"strings"
)

type Extractor struct {
	c *Config
}

func (e *Extractor) Load(c *Config) {
	e.c = c
}

func find(k []string, l int, d map[string]interface{}) interface{} {
	if l == 1 {
		return d[k[0]]
	}
	if c, ok := d[k[0]]; ok {
		if v, ok := c.(map[string]interface{}); ok {
			l--
			return find(k[1:], l, v)
		}
	}
	return nil
}

func (e *Extractor) Get(key string) interface{} {
	arr := strings.Split(key, ".")
	return find(arr, len(arr), e.c.data)
}

func (e *Extractor) String(key string) string {
	v := e.Get(key)
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func (e *Extractor) Int(key string) int {
	v := e.Get(key)
	if v == nil {
		return 0
	}
	if s, ok := v.(int); ok {
		return s
	}
	return 0
}

func (e *Extractor) Bool(key string) bool {
	v := e.Get(key)
	if v == nil {
		return false
	}
	if s, ok := v.(bool); ok {
		return s
	}
	return false
}

func (e *Extractor) Array(key string) []interface{} {
	v := e.Get(key)
	if v == nil {
		return nil
	}
	if array, ok := v.([]interface{}); ok {
		return array
	}
	return nil
}

func (e *Extractor) Map(key string) map[string]interface{} {
	v := e.Get(key)
	if v == nil {
		return nil
	}
	m := make(map[string]interface{})
	if mv, ok := v.(map[string]interface{}); ok {
		for k, item := range mv {
			m[k] = item
		}
		return m
	}
	return nil
}

func DefaultExtractor() IExtractor {
	return &Extractor{}
}