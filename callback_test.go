package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_Reload(t *testing.T) {
	v := 1
	c := Reload(func(i interface{}) {
		v = v + i.(int)
	})

	c.do(2)

	if !reflect.DeepEqual(v, 3) {
		t.Errorf("Test_Reload failed, expect 3, get %v ", v)
	}
}

func Test_Callback(t *testing.T) {
	path, _ := os.Getwd()
	file := path + "/test-config.yml"
	conf, _ := Init(Option{
		File:      file,
		Processor: YamlProcessor,
	})

	var v interface{}

	callback := Reload(func(i interface{}) {
		v = i
	})

	// get config value and register callback function
	v = conf.Get("test-config.user.name", callback, callback)
	if !reflect.DeepEqual(v, "calvin") {
		t.Errorf("Test_Callback before failed, expect %v, get %v ", "calvin", v)
	}

	// edit config file
	f, _ := ioutil.ReadFile(file)
	fs := string(f)
	nfs := strings.Replace(fs, "calvin", "ding", 1)
	err := ioutil.WriteFile(file, []byte(nfs), 0777)
	if err != nil {
		t.Error("file read error", err)
	}
	// rollback after test
	defer func() {
		err := ioutil.WriteFile(file, []byte(fs), 0777)
		if err != nil {
			t.Error("file read error", err)
		}
	}()

	// reload config file
	conf.Load(file)

	// check callback is effective
	if !reflect.DeepEqual(v, "ding") {
		t.Errorf("Test_Callback after failed, expect %v, get %v ", "ding", v)
	}
}
