package deployment

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	)

type Node struct {
	Name      string `yaml:"name"`
	Instances int `yaml:"instances"`
	Jobs       []Job  `yaml:"jobs"`
	Type      string `yaml:"lifecycle"`
}

type Job struct {
	Name      string `yaml:"name"`
}

type Deployment struct {
	Name  string   `yaml:"name"`
	Azs   []string `yaml:"azs"`
	Nodes []Node   `yaml:"instance_groups"`

	//Releases       interface{} `yaml:"releases"`
	//Stemcells      interface{} `yaml:"stemcells"`
	//Update         interface{} `yaml:"update"`
	//Variables      interface{} `yaml:"variables"`
}

func Load(file string) (Deployment, error) {

	deployment := Deployment{}

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return deployment, err
	}
	err = yaml.Unmarshal(yamlFile, &deployment)
	if err != nil {
		return deployment, err
	}

	return deployment, err
}
