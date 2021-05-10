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
package create_test

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"example.com/skeleton/pkg/create"
)

// Success/Failure chars
const (
	success = "\u2713"
	failure = "\u2717"
)

func setup(t *testing.T) (string, func()) {
	templatesDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unable to create temp directory %v", err)
	}

	tpl, err := ioutil.ReadFile("testdata/testfile.js.tpl")
	if err != nil {
		t.Fatalf("unable to read testdata file %v", err)
	}

	tempFile, err := ioutil.TempFile(templatesDir, "*.js.tpl")
	if err != nil {
		t.Fatalf("unable to create temp file %v", err)
	}
	defer tempFile.Close()

	if _, err := tempFile.Write(tpl); err != nil {
		t.Fatalf("unable to write to temp file %v", err)
	}

	teardown := func() {
		if err := os.RemoveAll(templatesDir); err != nil {
			t.Fatalf("unable to perform teardown %v", err)
		}
	}

	return templatesDir, teardown
}

func TestCreate(t *testing.T) {
	t.Log("Given the need to be able to scaffold a new project from templates")
	testID := 0

	templatesDir, teardown := setup(t)
	defer teardown()

	// =====================================================================
	// Creating output files based on templates in template directory
	t.Logf("\tTest %d:\tWhen given a directory containing at least one template.", testID)
	err := create.Create(templatesDir, templatesDir, "testdata/values.yaml")
	if err != nil {
		t.Fatalf("\t%s\tTest %d:\tShould be able to write files based on templates: %v", failure, testID, err)
	}
	t.Logf("\t%s\tTest %d:\tShould be able to write files based on templates.", success, testID)

	dirEntries, err := os.ReadDir(templatesDir)
	if err != nil {
		t.Fatal(err)
	}

	// =====================================================================
	// Validate output file extension
	if strings.Contains(dirEntries[0].Name(), ".tpl") {
		t.Fatalf("\t%s\tTest %d:\tShould remove \".tpl\" suffix: %v", failure, testID, dirEntries[0].Name())
	}
	t.Logf("\t%s\tTest %d:\tShould remove \".tpl\" suffix.", success, testID)

	data, err := ioutil.ReadFile(path.Join(templatesDir, dirEntries[0].Name()))
	if err != nil {
		t.Fatal(err)
	}

	// =====================================================================
	// Validate proper template interpolation
	if string(data) != "the-name" {
		t.Fatalf("\t%s\tTest %d:\tShould interpolate values into template: %v", failure, testID, err)
	}
	t.Logf("\t%s\tTest %d:\tShould interpolate values into template.", success, testID)
}
