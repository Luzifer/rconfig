package rconfig

import (
	"os"

	"go.yaml.in/yaml/v3"
)

// VarDefaultsFromYAMLFile reads contents of a file and calls VarDefaultsFromYAML
func VarDefaultsFromYAMLFile(filename string) map[string]string {
	data, err := os.ReadFile(filename) //#nosec:G304 // Loading file from var is intended
	if err != nil {
		return make(map[string]string)
	}

	return VarDefaultsFromYAML(data)
}

// VarDefaultsFromYAML creates a vardefaults map from YAML raw data
func VarDefaultsFromYAML(in []byte) map[string]string {
	out := make(map[string]string)
	err := yaml.Unmarshal(in, &out)
	if err != nil {
		return make(map[string]string)
	}
	return out
}
