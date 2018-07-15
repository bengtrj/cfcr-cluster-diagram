package main

import (
	"github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/generator/deployment"
	"github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/generator"
		"os"
	)

func main() {

	deployment, err := deployment.Load("/Users/bengthammarlund/go/src/github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/fixture/cfcr-v0.18.0-manifest.yml")
	if err != nil {
		panic("could not read the file")
	}

	generator := generator.Generator{
		Deployment:deployment,
	}

	generator.Generate(os.Stdout)

}
