/*
Copyright © 2024 Samuel Ireson <samuelireson@gmail.com>
*/
package compile

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var includeRegex = regexp.MustCompile(`\\include\{.*?\}`)

func findIncludes(courseDir string) []string {
	masterPath := filepath.Join(courseDir, "master.tex")
	masterFile, err := os.ReadFile(masterPath)
	if err != nil {
		log.Fatal(err)
	}
	content := string(masterFile)
	includeDirectives := includeRegex.FindAllString(content, -1)
	var chapters []string

	for _, includeDirective := range includeDirectives {
		chapter := strings.TrimPrefix(includeDirective, "\\include{")
		chapter = strings.TrimSuffix(chapter, "}")
		chapters = append(chapters, chapter)
	}

	return chapters
}

func compileMaster(courseDir string) {
	masterPath := filepath.Join(courseDir, "master")
	compileCommand := exec.Command(
		"latexmk",
		"-lualatex",
		"-cd",
		"-g",
		masterPath,
	)

	err := compileCommand.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s compiled successfuly", masterPath)

	outputDir := filepath.Join(courseDir, "output")
	os.Mkdir(outputDir, os.ModePerm)

	masterPDFPath := masterPath + ".pdf"
	outputPDFPath := filepath.Join(outputDir, "master.pdf")
	err = os.Rename(masterPDFPath, outputPDFPath)
	if err != nil {
		log.Fatal(err)
	}
}

func compileChapter(chapterPath, courseDir string) {
	useTeX := `-usepretex="\\includeonly{` + chapterPath + `}"`
	masterPath := filepath.Join(courseDir, "master.tex")
	compileCommand := exec.Command(
		"latexmk",
		"-lualatex",
		"-cd",
		"-g",
		useTeX,
		masterPath,
	)

	err := compileCommand.Run()
	if err != nil {
		log.Fatal(err)
	}

	outputDir := filepath.Join(courseDir, "output")
	os.Mkdir(outputDir, os.ModePerm)

	chapterName := strings.Split(chapterPath, "/")[1]
	outputPDFPath := filepath.Join(outputDir, chapterName+".pdf")
	masterPDFPath := strings.TrimSuffix(masterPath, filepath.Ext(masterPath)) + ".pdf"
	err = os.Rename(masterPDFPath, outputPDFPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s compiled successfuly", filepath.Join(courseDir, chapterPath))
}

func compileCourse(courseDir string) {
	chapters := findIncludes(courseDir)

	compileMaster(courseDir)

	for _, chapter := range chapters {
		compileChapter(chapter, courseDir)
	}
}

var CompileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile course notes from .tex to .pdf",
	Long: `Compile the course .tex files to .pdf, including separate files for
	chapters.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		compileCourse(args[0])
		if sync {
			outputPath := filepath.Join(args[0], "output")
			publicDownloadPath := filepath.Join("site/public", args[0])

			err := os.RemoveAll(publicDownloadPath)
			if err != nil {
				log.Fatal(err)
			}

			err = os.Rename(outputPath, publicDownloadPath)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Download path synced")
		}
	},
}

var sync bool

func init() {
	CompileCmd.Flags().BoolVarP(&sync, "sync", "s", false, "Sync the compiled files with those available for download")
}
