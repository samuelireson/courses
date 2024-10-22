/*
Copyright © 2024 Samuel Ireson <samuelireson@gmail.com>
*/
package compile

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

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
	chaptersDir := filepath.Join(courseDir, "chapters")
	chapters, err := os.ReadDir(chaptersDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, chapter := range chapters {
		chapterName := chapter.Name()
		chapterFilepath := strings.TrimSuffix(chapterName, filepath.Ext(chapterName))
		chapterPath := filepath.Join("chapters", chapterFilepath)
		compileChapter(chapterPath, courseDir)
	}

	compileMaster(courseDir)
}

var CompileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile course notes from .tex to .pdf",
	Long: `Compile the course .tex files to .pdf, including separate files for
	chapters.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		compileCourse(args[0])
	},
}

func init() {
}
