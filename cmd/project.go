package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/CurtisMIT/vulcan/tpl"
)

// Project struct for any boilerplate
type Project struct {
	AbsolutePath string
	FileName     string
	Tpl          string
	Db           string
}

// Create creates a directory
func (p *Project) Create() error {
	projFile, err := os.Create(fmt.Sprintf("%s/%s", p.AbsolutePath, p.FileName))
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
