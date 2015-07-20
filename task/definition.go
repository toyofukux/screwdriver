package task

import (
	"bytes"
	"text/template"

	"github.com/aws/aws-sdk-go/service/ecs"

	"gopkg.in/yaml.v2"
)

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
func BindYml(b []byte) ([]ecs.ContainerDefinition, error) {
	yml := make(map[string]ecs.ContainerDefinition)
	err := yaml.Unmarshal(b, &yml)
	if err != nil {
		return nil, err
	}
	var containers []ecs.ContainerDefinition
	for k, v := range yml {
		v.Name = &k
		containers = append(containers, v)
	}
	return containers, nil
}
