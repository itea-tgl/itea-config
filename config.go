package config

import (
	"errors"
	"path"
	"strings"
	"sync"
)

type IProcessor interface {
	Load(string) (map[string]interface{}, error)
}

type IExtractor interface {
	Load(c *Config)
	Get(string) interface{}
	String(string) string
	Int(string) int
	Bool(string) bool
	Array(string) []interface{}
	Map(string) map[string]interface{}
}

// ProcessorConstruct is the type for a function capable of constructing new IProcessor.
type ProcessorConstruct func() IProcessor

// ExtractorConstruct is the type for a function capable of constructing new IExtractor.
type ExtractorConstruct func() IExtractor

// Option is used to pass multiple configuration options to Config's constructors.
type Option struct {
	File 		string
	Processor 	ProcessorConstruct
	Extractor	ExtractorConstruct
}

type Config struct {
	data map[string]interface{}
	processor IProcessor
	extractor IExtractor
	l sync.RWMutex
}

func Init(option Option) (c *Config, err error) {
	if option.Processor == nil {
		return nil, errors.New("undefined processor")
	}

	c = &Config{
		data: make(map[string]interface{}),
		processor: option.Processor(),
		extractor: DefaultExtractor(),
	}

	if option.Extractor != nil {
		c.setExtractor(option.Extractor())
	}

	c.extractor.Load(c)

	if option.File != "" {
		err = c.Load(option.File)
	}

	return c, err
}

func (c *Config) setProcessor(p IProcessor) {
	c.processor = p
}

func (c *Config) setExtractor(e IExtractor) {
	c.extractor = e
}

func (c *Config) setData(k string, v interface{}) {
	c.data[k] = v
}

func (c *Config) Load(file string) (e error) {
	d, e := c.processor.Load(file)
	if e != nil {
		return
	}
	c.l.Lock()
	defer c.l.Unlock()
	c.setData(filename(file), d)
	return
}

func (c *Config) Get(key string) interface{} {
	return c.extractor.Get(key)
}

func (c *Config) Int(key string) int {
	return c.extractor.Int(key)
}

func (c *Config) String(key string) string {
	return c.extractor.String(key)
}

func (c *Config) Bool(key string) bool {
	return c.extractor.Bool(key)
}

func (c *Config) Array(key string) []interface{} {
	return c.extractor.Array(key)
}

func (c *Config) Map(key string) map[string]interface{} {
	return c.extractor.Map(key)
}

func filename(file string) string {
	filename := path.Base(file)
	suffix := path.Ext(filename)
	return strings.TrimSuffix(filename, suffix)
}