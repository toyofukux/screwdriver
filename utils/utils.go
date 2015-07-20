package utils

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go/service/ecs"
	"gopkg.in/yaml.v2"
)

const (
	// EnvPrefix is the prefix to load environment variable.
	EnvPrefix = "SCREW_"
)

// LoadScrewEnvs is laod environment variable and return list of environment.
func LoadScrewEnvs() map[string]string {
	data := os.Environ()
	items := make(map[string]string)
	for _, val := range data {
		splits := strings.SplitN(val, "=", 2)
		key := splits[0]
		if strings.HasPrefix(key, EnvPrefix) {
			k := strings.TrimPrefix(key, EnvPrefix)
			items[k] = splits[1]
		}
	}
	return items
}

// ExpandTemplate is expand placeholder in yml by environment variables.
func ExpandTemplate(b []byte, envs map[string]string) ([]byte, error) {
	tmpl, err := template.New("task.yml").Parse(string(b))
	if err != nil {
		return nil, err
	}
	var expanded bytes.Buffer
	err = tmpl.Execute(&expanded, envs)
	if err != nil {
		return nil, err
	}
	e := expanded.Bytes()
	return e, nil
}

// BindYml bind yml to ecs.ContainerDefinition struct array.
func BindYml(b []byte) ([]*ecs.ContainerDefinition, error) {
	yml := make(map[string]ecs.ContainerDefinition)
	err := yaml.Unmarshal(b, &yml)
	if err != nil {
		return nil, err
	}
	var containers []*ecs.ContainerDefinition
	for k, v := range yml {
		// for
		key, container := k, v
		container.Name = &key
		containers = append(containers, &container)
	}
	return containers, nil
}
