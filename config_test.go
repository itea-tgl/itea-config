package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_Yaml(t *testing.T) {
	path, _ := os.Getwd()

	conf, _ := Init(Option{
		File:      path + "/test-config.yml",
		Processor: YamlProcessor,
	})

	type Struct struct {
		Param1 string
		Param2 string
	}

	sa := conf.StructArray("test-config.struct_array", Struct{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			"calvin",
			conf.Get("test-config.user.name"),
		},
		{
			"calvin",
			conf.String("test-config.user.name"),
		},
		{
			10,
			conf.Int("test-config.user.age"),
		},
		{
			true,
			conf.Bool("test-config.user.male"),
		},
		{
			false,
			conf.Bool("test-config.user.female"),
		},
		{
			[]interface{}{"aaa", "bbb", "ccc"},
			conf.Array("test-config.user.list"),
		},
		{
			map[string]interface{}{"aa":"aa", "bb":"bb", "11":11},
			conf.Map("test-config.user.property"),
		},
		{
			&Struct{Param1: "aaa", Param2: "bbb",},
			conf.Struct("test-config.struct", Struct{}),
		},
		{
			&Struct{Param1: "aaa", Param2: "bbb",},
			sa[0],
		},
		{
			&Struct{Param1: "ccc", Param2: "ddd",},
			sa[1],
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Test_Yaml %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}

}

func Test_Json(t *testing.T) {
	path, _ := os.Getwd()

	conf, _ := Init(Option{
		File:      path + "/test-config.json",
		Processor: JsonProcessor,
	})

	type Struct struct {
		Param1 string
		Param2 string
	}

	sa := conf.StructArray("test-config.struct_array", Struct{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			"calvin",
			conf.Get("test-config.user.name"),
		},
		{
			"calvin",
			conf.String("test-config.user.name"),
		},
		{
			10,
			conf.Int("test-config.user.age"),
		},
		{
			true,
			conf.Bool("test-config.user.male"),
		},
		{
			false,
			conf.Bool("test-config.user.female"),
		},
		{
			[]interface{}{"aaa", "bbb", "ccc"},
			conf.Array("test-config.user.list"),
		},
		{
			map[string]interface{}{"aa":"aa", "bb":"bb", "11":11},
			conf.Map("test-config.user.property"),
		},
		{
			&Struct{Param1: "aaa", Param2: "bbb",},
			conf.Struct("test-config.struct", Struct{}),
		},
		{
			&Struct{Param1: "aaa", Param2: "bbb",},
			sa[0],
		},
		{
			&Struct{Param1: "ccc", Param2: "ddd",},
			sa[1],
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Test_Yaml %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}

}

func Test_Ini(t *testing.T) {
	path, _ := os.Getwd()

	conf, _ := Init(Option{
		File:      path + "/test-config.ini",
		Processor: IniProcessor,
	})

	type Struct struct {
		Param1 string
		Param2 string
	}

	//sa := conf.StructArray("test-config.struct_array", Struct{})

	tests := []struct {
		want interface{}
		result interface{}
	} {
		{
			"calvin",
			conf.Get("test-config.user.name"),
		},
		{
			"aa",
			conf.String("test-config.user.property.aa"),
		},
		{
			10,
			conf.Int("test-config.user.age"),
		},
		{
			11,
			conf.Int("test-config.user.property.11"),
		},
		{
			true,
			conf.Bool("test-config.user.male"),
		},
		{
			false,
			conf.Bool("test-config.user.female"),
		},
		{
			map[string]interface{}{"aa":"aa", "bb":"bb", "11":"11"},
			conf.Map("test-config.user.property"),
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.want, test.result) {
			t.Errorf("Test_Yaml %d failed, expect %v, get %v ", i, test.want, test.result)
		}
	}

}