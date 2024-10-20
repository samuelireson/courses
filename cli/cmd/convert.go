/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func generateOutputFilePath(input string) string {
	if customOutputPath != "" {
		return customOutputPath
	} else {
		inputDir, inputFileName := filepath.Split(input)
		baseName := strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName))
		outputFileName := baseName + ".mdx"
		outputDir := strings.Replace(inputDir, "notes", "site/src/content/docs", 1)
		outputDir = strings.Replace(outputDir, "chapters/", "", 1)
		outputPath := filepath.Join(outputDir, outputFileName)
		os.MkdirAll(outputDir, os.ModePerm)
		return outputPath
	}
}

func processFile(inputPath string) {
	outputPath := generateOutputFilePath(inputPath)
	fi, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	fo := convertTeXToMDX(fi)

	err = os.WriteFile(outputPath, fo, 0644)
	if err != nil {
		panic(err)
	}
}

func processDir(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		inputPath := filepath.Join(dirPath, file.Name())
		processFile(inputPath)
	}
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert course notes from .tex to .mdx",
	Long: `Convert LaTeX notes for a course to MarkdownX format, which can be
	rendered on the web.`,
	// Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if customBibPath != "" {
			parseBibliography(customBibPath)
			os.Exit(0)
		}

		dirPath := filepath.Join(args[0], "/chapters")
		processDir(dirPath)
		if continuous {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("Watching for changes to %s\n", dirPath)
			}
			defer watcher.Close()

			done := make(chan bool)
			timers := make(map[string]*time.Timer)

			go func() {
				for {
					select {
					case event := <-watcher.Events:
						if event.Op&fsnotify.Write == fsnotify.Write {
							if timer, exists := timers[event.Name]; exists {
								timer.Stop()
							}

							timers[event.Name] = time.AfterFunc(1*time.Second, func() {
								fmt.Println("Files changed, re-converting")
								processDir(dirPath)
								delete(timers, event.Name)
							})
						}
					case err := <-watcher.Errors:
						panic(err)
					}
				}
			}()

			err = watcher.Add(dirPath)
			if err != nil {
				panic(err)
			}

			<-done
		}
	},
}

var continuous bool
var customOutputPath string
var customBibPath string

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Watch and continuously convert")
	convertCmd.Flags().StringVarP(&customOutputPath, "output", "o", "", "Specify a custom output path")
	convertCmd.Flags().StringVarP(&customBibPath, "bib", "b", "", "Specify a custom bibliography path")
}
