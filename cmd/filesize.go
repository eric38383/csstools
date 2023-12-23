package cmd

import (
	"fmt"

	"github.com/eric38383/csstools/utils"
	"github.com/spf13/cobra"
)

var ext string

func init() {
	rootCmd.AddCommand(fileSizeCmd)
	fileSizeCmd.Flags().StringVarP(&ext, "ext", "", "", "The File Extension")
	fileSizeCmd.MarkFlagRequired("ext")
}

var fileSizeCmd = &cobra.Command{
	Use:   "filesize",
	Short: "Get size of each file by extension and total size of directory by extension",
	Run: func(cmd *cobra.Command, args []string) {
		ext, _ := cmd.Flags().GetString("ext")
		if ext != ".scss" {
			fmt.Println("This process is currently only setup to check SASS files. Please use .scss")
			return
		}
		output := utils.GetFileSizeByExt(ext)
		for k, v := range output {
			fmt.Printf("%s: %.2fkb, %.2fmb \n", k, v.TotalKB, v.TotalMB)
			for _, file := range v.Files {
				fmt.Printf("%s: %.2fkb, %.2fmb \n", file.Name, file.KB, file.MB)
			}
		}
	},
}
