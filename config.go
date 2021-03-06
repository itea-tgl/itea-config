package config

import (
	"errors"
	"sync"
)

type IProcessor interface {
	Load(string) (map[string]interface{}, error)
}

type IExtractor interface {
	Load(interface{}, string)
	Get(string, ...Callback) interface{}
	String(string, ...Callback) string
	Int(string, ...Callback) int
	Bool(string, ...Callback) bool
	Array(string, ...Callback) []interface{}
	Map(string, ...Callback) map[string]interface{}
	Struct(string, interface{}, ...Callback) interface{}
	StructArray(string, interface{}, ...Callback) []interface{}
}

// ProcessorConstruct is the type for a function capable of constructing new IProcessor.
type ProcessorConstruct func() IProcessor

// ExtractorConstruct is the type for a function capable of constructing new IExtractor.
type ExtractorConstruct func() IExtractor

// Option is used to pass multiple configuration options to Config's constructors.
type Option struct {
	File      string
	Processor ProcessorConstruct
	Extractor ExtractorConstruct
}

type Config struct {
	data      map[string]interface{}
	processor IProcessor
	extractor IExtractor
	l         sync.RWMutex
}

func Init(option Option) (c *Config, err error) {
	if option.Processor == nil {
		return nil, errors.New("undefined processor")
	}

	c = &Config{
		processor: option.Processor(),
		extractor: DefaultExtractor(),
	}

	if option.Extractor != nil {
		c.extractor = option.Extractor()
	}

	if option.File != "" {
		err = c.Load(option.File)
	}

	return c, err
}

func (c *Config) Load(file string) (e error) {
	c.l.Lock()
	defer c.l.Unlock()
	return Loader(func() string {
		return file
	}).load(c.processor, c.extractor)
}

func (c *Config) Get(key string, b ...Callback) interface{} {
	return c.extractor.Get(key, b...)
}

func (c *Config) Int(key string, b ...Callback) int {
	return c.extractor.Int(key, b...)
}

func (c *Config) String(key string, b ...Callback) string {
	return c.extractor.String(key, b...)
}

func (c *Config) Bool(key string, b ...Callback) bool {
	return c.extractor.Bool(key, b...)
}

func (c *Config) Array(key string, b ...Callback) []interface{} {
	return c.extractor.Array(key, b...)
}

func (c *Config) Map(key string, b ...Callback) map[string]interface{} {
	return c.extractor.Map(key, b...)
}

func (c *Config) Struct(key string, s interface{}, b ...Callback) interface{} {
	return c.extractor.Struct(key, s, b...)
}

func (c *Config) StructArray(key string, s interface{}, b ...Callback) []interface{} {
	return c.extractor.StructArray(key, s, b...)
}
