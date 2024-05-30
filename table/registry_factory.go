package table

import (
	"errors"
)

var registries = make(map[string]Table)

func Register(tp Table, pluginName string) error {
	if _, ok := registries[pluginName]; ok {
		return errors.New("Cannot add duplicate registry: " + pluginName)
	}
	registries[pluginName] = tp
	return nil
}

func GetAll() map[string]Table {
	return registries
}

func Get(tableName string) Table {
	return registries[tableName]
}
