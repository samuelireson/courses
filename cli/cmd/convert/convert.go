/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package convert

import (
	"log"
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

func generateBibPath(input string) string {
	if customBibPath != "" {
		return customBibPath
	}
	return filepath.Join(input, "bibliography.bib")
}

func processFile(bibliography bibliography, inputPath string) {
	outputPath := generateOutputFilePath(inputPath)
	fi, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	content := string(fi)
	content = convertTeXToMDX(content)
	content = convertCitationsToFootnotes(bibliography, content)
	content = addDownloadLinks(content, inputPath)
	fo := []byte(content)

	err = os.WriteFile(outputPath, fo, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s converted successfuly", inputPath)
}

func processDir(bibliography bibliography, dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".tex" {
			inputPath := filepath.Join(dirPath, file.Name())
			processFile(bibliography, inputPath)
		}
	}
}

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert course notes from .tex to .mdx",
	Long: `Convert LaTeX notes for a course to MarkdownX format, which can be
	rendered on the web.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bibPath := generateBibPath(args[0])
		bibliography := parseBibliography(bibPath)

		dirPath := filepath.Join(args[0], "/chapters")
		processDir(bibliography, dirPath)
		if continuous {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Watching for changes to %s\n", dirPath)
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
								log.Println("Files changed, re-converting")
								processDir(bibliography, dirPath)
								delete(timers, event.Name)
							})
						}
					case err := <-watcher.Errors:
						log.Fatal(err)
					}
				}
			}()

			err = watcher.Add(dirPath)
			if err != nil {
				log.Fatal(err)
			}

			<-done
		}
	},
}

var continuous bool
var customOutputPath string
var customBibPath string

func init() {
	ConvertCmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Watch and continuously convert")
	ConvertCmd.Flags().StringVarP(&customOutputPath, "output", "o", "", "Specify a custom output path")
	ConvertCmd.Flags().StringVarP(&customBibPath, "bib", "b", "", "Specify a custom bibliography path")
}
