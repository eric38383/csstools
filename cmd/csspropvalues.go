package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/eric38383/csstools/parser"
	"github.com/spf13/cobra"
)

var filepath string

func init() {
	rootCmd.AddCommand(cssPropValuesCmd)
	cssPropValuesCmd.Flags().StringVarP(&filepath, "filepath", "", "", "The CSS Filepath")
}

var skippedProperties = []string{"@font-face"}

var cssPropValuesCmd = &cobra.Command{
	Use:   "csspropvalues",
	Short: "Get CSS Properties and their unique values",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("filepath")
		file, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := make(map[string][]string)
		str := string(file)
		rules := parser.New(str).Stylesheet()
		for _, rule := range rules {
			if slices.Contains(skippedProperties, rule.Name) {
				continue
			}

			for _, declaration := range rule.Declarations {
				var prop = declaration.Property
				if slices.Contains(data[prop], declaration.Value) {
					continue
				}

				if values, ok := data[prop]; ok {
					data[prop] = append(values, declaration.Value)
				} else {
					data[prop] = []string{declaration.Value}
				}
			}
		}

		for key, value := range data {
			fmt.Printf("%s: %v\n", key, value)
		}
	},
}
