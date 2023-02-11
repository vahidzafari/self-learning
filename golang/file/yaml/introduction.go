package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// The Go standard library does not include support for YAML files, which means that you
// should look at external packages for YAML support. There exist three main packages
// that allow you to work with YAML from Go:
// - https://github.com/kylelemons/go-gypsy
// - https://github.com/go-yaml/yaml
// - https://github.com/goccy/go-yaml

var yamlfile = `
image: Golang
matrix:
	docker: python
	version: [2.7, 3.9]
`

type Mat struct {
	DockerImage string    `yaml:"docker"`
	Version     []float32 `yaml:",flow"`
}

type YAML struct {
	Image  string
	Matrix Mat
}

func main() {
	data := YAML{}

	err := yaml.Unmarshal([]byte(yamlfile), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After Unmarshal (Structure):\n%v\n\n", data)

	d, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After Marshal (YAML code):\n%s\n", string(d))
}
