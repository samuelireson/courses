/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package cmd

import (
	"courses/cmd/compile"
	"courses/cmd/convert"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "courses",
	Short: "A CLI to manage course notes.",
	Long: `courses is a CLI which lets you convert course notes between formats,
	quickly set up new courses, and write beautiful notes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(convert.ConvertCmd)
	rootCmd.AddCommand(compile.CompileCmd)
}
