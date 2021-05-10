/*
Copyright © 2021 Shane Unger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package create

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type val map[string]interface{}

// Create reads the templates in templatesDir, executes them and writes them to outputDir
func Create(templatesDir, outputDir, valuesFilepath string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("Create : creating skeleton failed %v", r)
		}
	}()

	bs, err := ioutil.ReadFile(valuesFilepath)
	if err != nil {
		return errors.Wrap(err, "Create : could not read values file")
	}

	values := make(val)
	if err := yaml.Unmarshal(bs, &values); err != nil {
		return errors.Wrap(err, "Create : unable to parse values.yaml")
	}

	walk := func(p string, d fs.FileInfo, e error) error {
		if d.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(p)
		if err != nil {
			return errors.Wrapf(err, "Create : error reading file %q", d.Name())
		}

		// If file is a template execute the template and write the file contents in output file
		if strings.Contains(p, ".tpl") {
			result, err := executeTemplate(d.Name(), data, values)
			if err != nil {
				return errors.Wrap(err, "Create: unable to execute template")
			}

			// remove `.tpl` suffix from path
			n := strings.TrimSuffix(p, ".tpl")
			return writeFile(n, templatesDir, outputDir, result)
		}

		// File wasn't a template, simply copy the contents to the new file
		return writeFile(p, templatesDir, outputDir, data)
	}

	// walk templates directory to map contents and structure to output dir
	if err := filepath.Walk(templatesDir, walk); err != nil {
		return errors.Wrap(err, "Create : error creating directory from template")
	}

	return nil
}

// executeTemlate takes the template name, the template contents and the values to interpolate
// and returns a byte slice representing the executed template
func executeTemplate(name string, fileContents []byte, values val) ([]byte, error) {
	tpl, err := template.New(name).Funcs(sprig.TxtFuncMap()).Parse(string(fileContents))
	if err != nil {
		return nil, errors.Wrap(err, "executeTemplate : error parsing template")
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, values); err != nil {
		return nil, errors.Wrap(err, "executeTemplate : error executing template")
	}

	return buf.Bytes(), nil
}

func writeFile(templateFilePath, templatesDir, outputDir string, data []byte) error {
	// replace templates dir path with the user requested output dir path
	writePath := strings.Replace(templateFilePath, templatesDir, outputDir, 1)

	writeDir := filepath.Dir(writePath)
	if _, err := os.Stat(writeDir); os.IsNotExist(err) {
		if err := os.MkdirAll(writeDir, os.ModePerm); err != nil {
			return errors.Wrapf(err, "Create : unable to create dir %q", writeDir)
		}
	}

	return os.WriteFile(writePath, data, os.ModePerm)
}
