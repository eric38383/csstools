package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fileSizeCmd)
}

var fileSizeCmd = &cobra.Command{
	Use:   "filesize",
	Short: "I will tell the file of each file and directory by filter by an extension",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here We go!")
	},
}
