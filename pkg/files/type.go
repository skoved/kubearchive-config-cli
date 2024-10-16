// Copyright skoved

package files

import (
	"os"

	"gopkg.in/yaml.v3"
)

func IsYaml(file string) bool {
	reader, err := os.Open(file)
	if err != nil {
		return false
	}
	dec := yaml.NewDecoder(reader)
	var data map[string]interface{}
	err = dec.Decode(&data)
	if err != nil {
		return false
	}
	return true
}
