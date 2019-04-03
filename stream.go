package gompeg

import (
	"github.com/fatih/structs"
	"reflect"
	"strings"
)

func (m *Media) Build() []string {
	var commands []string
	fields := structs.Names(m)
	for _, v := range fields {
		value := reflect.ValueOf(m).MethodByName(strings.Title(v))
		if (value != reflect.Value{}) {
			result := value.Call([]reflect.Value{})
			if v, ok := result[0].Interface().([]string); ok {
				commands = append(commands, v...)
			}
		}
	}
	return commands
}
