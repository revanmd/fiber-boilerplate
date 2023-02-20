package config

import (
	"reflect"
)

var (
	Cfg *Config
)

type Config struct {
	ServerPort       string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

func (c Config) getAttrName() []string {
	t := reflect.TypeOf(c)
	names := make([]string, t.NumField())
	for i := range names {
		names[i] = t.Field(i).Name
	}
	return names
}
func (c *Config) setProperty(propName string, propValue string) *Config {
	reflect.ValueOf(c).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return c
}

func (c *Config) getConfiguration() *Config {
	return c
}
