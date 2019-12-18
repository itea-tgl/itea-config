package config

import "io/ioutil"

type Base struct {

}

func (Base) load(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
