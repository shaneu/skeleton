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
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"example.com/skeleton/pkg/create"
	"github.com/spf13/cobra"
)

var valuesFilepath string

var createCmd = &cobra.Command{
	Use:   "create [flags] TEMPLATES_DIR OUTPUT_DIR",
	Short: "create output directory structure from templates directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("incorrect number of arguments")
		}

		templatesDir := filepath.Clean(args[0])
		outputDir := args[1]

		err := create.Create(templatesDir, outputDir, valuesFilepath)
		if err != nil {
			_ = os.RemoveAll(outputDir)
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&valuesFilepath, "values", "v", "", "path to values.yaml file")
}
