package cmd

import (
	"fmt"
	"os"
	"text/template"
	"vulcan/tpl"
)

// Project struct for any boilerplate
type Project struct {
	AbsolutePath string
}

// Create creates a directory
func (p *Project) Create() error {
	projFile, err := os.Create(fmt.Sprintf("%s/index.js", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer projFile.Close()

	projTemplate := template.Must(template.New("sub").Parse(string(tpl.AddNodeTemplate())))
	err = projTemplate.Execute(projFile, p)
	if err != nil {
		return err
	}
	return nil
}
