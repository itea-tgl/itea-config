package config

type Callback interface {
	do(interface{})
}

type Reload func(v interface{})

func (r Reload) do(v interface{}) {
	r(v)
}