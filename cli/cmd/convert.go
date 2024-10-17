/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package cmd

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type regexPattern struct {
	captureGroup *regexp.Regexp
	replacement  string
}

var header = `---
title: $1
---
import { Aside } from '\@components';
import { Tabs, TabItem } from '\@astrojs/starlight/components';
`

var basicRegexPatterns = []regexPattern{
	// document organisation
	{regexp.MustCompile(`\\begin\{chout\}`), "<div style='text-align: center'><em>"},
	{regexp.MustCompile(`\\end\{chout\}`), "</em></div>"},
	{regexp.MustCompile(`\\chapter\{(.*?)\}`), header},
	{regexp.MustCompile(`\\section\{(.*?)\}`), "## $1"},
	{regexp.MustCompile(`\\subsection\{(.*?)\}`), "### $1"},

	// theorem environments
	{regexp.MustCompile(`\\begin\{definition\}`), "<Aside type='definition' title='Definition' >"},
	{regexp.MustCompile(`\\begin\{(theorem|lemma|proposition|corollary)\}`), "<Aside type='result' title='" + strings.ToUpper("$1") + "' >"},
	{regexp.MustCompile(`\\begin\{(example|nonexample)\}`), "<Aside type='result' title='" + strings.ToUpper("$1") + "' >"},
	{regexp.MustCompile(`\\begin\{(notation|remark)\}`), "<Aside type='result' title='" + strings.ToUpper("$1") + "' >"},
	{regexp.MustCompile(`\\end\{(definition|theorem|lemma|proposition|corollary|example|nonexample|notation|remark)\}`), "</Aside>"},

	// maths environments
	{regexp.MustCompile(`(\\begin\{align\*\})`), "$$$$\n$1"},
	{regexp.MustCompile(`(\\end\{align\*\})`), "$1\n$$$$"},

	// exercises
	{regexp.MustCompile(`\\begin\{exercise\}`), "<Tabs>"},
	{regexp.MustCompile(`\\end\{exercise\}`), "</Tabs>"},
	{regexp.MustCompile(`\\begin\{problem\}`), "<TabItem label='Problem'>"},
	{regexp.MustCompile(`\\begin\{solution\}`), "<TabItem label='Solution'>"},
	{regexp.MustCompile(`\\end\{(problem|solution)\}`), "</TabItem>"},

	// badges
	{regexp.MustCompile(`\\basic`), ":badge[Basic]{variant=success}"},
	{regexp.MustCompile(`\\intermediate`), ":badge[Intermediate]{variant=warning}"},
	{regexp.MustCompile(`\\challenging`), ":badge[Challenging]{variant=danger}"},
}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert course notes from .tex to .mdx",
	Long: `Convert LaTeX notes for a course to MarkdownX format, which can be
	rendered on the web.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		outputPath := generateOutputFilePath(inputPath)
		convertTeXtoMDX(inputPath, outputPath)
	},
}

func generateOutputFilePath(input string) string {
	inputDir, inputFileName := filepath.Split(input)
	baseName := strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName))
	outputFileName := baseName + ".mdx"
	outputDir := strings.Replace(inputDir, "notes", "site/src/content/docs", 1)
	outputDir = strings.Replace(outputDir, "chapters/", "", 1)
	outputPath := filepath.Join(outputDir, outputFileName)
	os.MkdirAll(outputDir, os.ModePerm)
	return outputPath
}

func convertTeXtoMDX(inputPath string, outputPath string) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	for _, element := range basicRegexPatterns {
		content = element.captureGroup.ReplaceAll(content, []byte(element.replacement))
	}

	err = os.WriteFile(outputPath, content, 0644)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
