package config

import (
	"fmt"
	"reflect"
	"strings"
)

type Value struct {
	Key 		string
	Type 		string
	Value		interface{}
	Child 		map[string]*Value
	Callback 	[]Callback
}

func (v *Value) Get(key string, b ...Callback) interface{} {
	d := v.value(key)
	if d == nil {
		return nil
	}

	go callback(d, b...)

	return d.Value
}

func (v *Value) String(key string, b ...Callback) string {
	i := v.Get(key, b...)
	if i == nil {
		return ""
	}
	if s, ok := i.(string); ok {
		return s
	}
	return ""
}

func (v *Value) Int(key string, b ...Callback) int {
	i := v.Get(key, b...)
	if i == nil {
		return 0
	}
	if s, ok := i.(int); ok {
		return s
	}
	return 0
}

func (v *Value) Bool(key string, b ...Callback) bool {
	i := v.Get(key, b...)
	if i == nil {
		return false
	}
	if s, ok := i.(bool); ok {
		return s
	}
	return false
}

func (v *Value) Array(key string, b ...Callback) []interface{} {
	i := v.Get(key, b...)
	if i == nil {
		return nil
	}
	if a, ok := i.([]interface{}); ok {
		return a
	}
	return nil
}

func (v *Value) Map(key string, b ...Callback) map[string]interface{} {
	i := v.Get(key, b...)
	if i == nil {
		return nil
	}
	m := make(map[string]interface{})
	if mv, ok := i.(map[string]interface{}); ok {
		for k, item := range mv {
			m[k] = item
		}
		return m
	}
	return nil
}

func (v *Value) Load(m map[string]interface{}, name string) {
	v.recursionValue(v, name, m)
}

func (v *Value) value(key string) *Value {
	keyArr := strings.Split(key, ".")
	l := len(keyArr)
	if l == 0 {
		return nil
	}
	if c, ok := v.Child[keyArr[0]]; ok {
		if l == 1 {
			return c
		}
		return c.value(key[len(keyArr[0])+1:])
	}
	return nil
}

func (v *Value) recursionValue(parent *Value, k string, i interface{}) {
	var value *Value
	key := k

	if parent.Key != "" {
		key = fmt.Sprintf("%s.%s", parent.Key, key)
	}

	var t string
	switch i.(type) {
	case int:
		t = "int"
		break
	case string:
		t = "string"
		break
	case bool:
		t = "bool"
		break
	case map[string]interface{}:
		t = "map"
		break
	}

	if d, ok := parent.Child[k]; ok {
		value = d
		if !reflect.DeepEqual(d.Value, i) {
			defer trigger(d)
		}
		value.Value = i
	} else {
		value = &Value{
			Key:      key,
			Type:     t,
			Value:	  i,
			Child:    make(map[string]*Value),
			Callback: make([]Callback, 0),
		}
	}
	
	if t == "map" {
		for k1, v1 := range i.(map[string]interface{}) {
			v.recursionValue(value, k1, v1)
		}
	}

	parent.Child[k] = value
}

func trigger(value *Value) {
	for _, c := range value.Callback {
		c.do(value.Value)
	}
}

func callback(value *Value, b ...Callback) {
	value.Callback = append(value.Callback, func(b []Callback) []Callback {
		var callback []Callback
		for _, c := range b {
			if func(b Callback) bool {
				for _, c := range value.Callback {
					if reflect.ValueOf(b) == reflect.ValueOf(c) {
						return false
					}
				}
				return true
			}(c) {
				callback = append(callback, c)
			}
		}
		return callback
	}(b)...)
}

func DefaultExtractor() *Value {
	return &Value{
		Key:      "",
		Type:     "map",
		Value:    nil,
		Child:    make(map[string]*Value),
		Callback: make([]Callback, 0),
	}
}