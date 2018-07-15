package generator

import (
	"github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/generator/deployment"
		"io"
		"text/template"
	"github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/generator/cluster"
	"fmt"
	"bytes"
	"io/ioutil"
	)

type Generator struct {
	Deployment deployment.Deployment
}

func (g Generator) Generate(writer io.Writer) {

	tmp := template.Must(template.ParseFiles("/Users/bengthammarlund/go/src/github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/templates/cluster.tmpl"))

	c, err := cluster.NewCluster(g.Deployment)

	if err != nil {
		panic(fmt.Sprintf("Could not get cluster %v", err.Error()))
	}

	buf := bytes.NewBufferString("")

	tmp.Execute(buf, c)

	code := template.Must(template.ParseFiles("/Users/bengthammarlund/go/src/github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/templates/code.js.tmpl"))

	finalFile := bytes.NewBufferString("")

	code.Execute(finalFile, buf.String())

	ioutil.WriteFile("/Users/bengthammarlund/workspace/cfcr-diagram-playground/architecture/code.js", finalFile.Bytes(), 0644)
}
