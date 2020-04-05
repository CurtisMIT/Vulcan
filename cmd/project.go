package cmd

import (
	"fmt"
	"github.com/CurtisMIT/vulcan/tpl"
	"os"
	"text/template"
)

// Project struct for any boilerplate
type Project struct {
	AbsolutePath string
	FileType     string
}

// Create creates a directory
func (p *Project) Create() error {
	projFile, err := os.Create(fmt.Sprintf("%s/%s", p.AbsolutePath, p.FileType))
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
