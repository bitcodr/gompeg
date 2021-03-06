package gompeg

import (
	"github.com/fatih/structs"
	"os/exec"
	"reflect"
	"strings"
)

func (m *Media) Build() error {
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
	err := exec.Command("ffmpeg", commands...).Run()
	if err != nil {
		return err
	}
	return nil
}
