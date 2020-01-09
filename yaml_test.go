package config

import (
	"os"
	"testing"
)

func Test_Yaml_Load(t *testing.T) {
	path, _ := os.Getwd()

	yaml := YamlProcessor()
	_, err := yaml.Load(path + "/test-config.yml")
	if err != nil {
		t.Errorf("Test_Yaml_Load failed, %s", err.Error())
	}
}