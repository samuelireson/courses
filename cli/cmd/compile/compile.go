/*
Copyright © 2024 Samuel Ireson <samuelireson@gmail.com>
*/
package compile

import (
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
	log.Printf("%s compiled successfully", masterPath)

	outputDir := filepath.Join(courseDir, "output")
	os.Mkdir(outputDir, os.ModePerm)

	masterPDFPath := masterPath + ".pdf"
	outputPDFPath := filepath.Join(outputDir, "master.pdf")
	err = os.Rename(masterPDFPath, outputPDFPath)
	if err != nil {
		log.Fatal(err)
	}
}

func compileChapter(chapter, courseDir string) {
	includeChapterPath := filepath.Join("chapters", chapter)
	chapterPath := filepath.Join(courseDir, includeChapterPath)
	useTeX := `-usepretex="\\includeonly{` + includeChapterPath + `}"`
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

	chapterName := strings.TrimSuffix(chapter, filepath.Ext(chapter))
	outputPDFPath := filepath.Join(outputDir, chapterName+".pdf")
	masterPDFPath := strings.TrimSuffix(masterPath, filepath.Ext(masterPath)) + ".pdf"
	err = os.Rename(masterPDFPath, outputPDFPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s compiled successfully", filepath.Join(courseDir, chapterPath))
}

func compileCourse(courseDir string) {
	chapterDir := filepath.Join(courseDir, "chapters")
	chapters, err := os.ReadDir(chapterDir)
	if err != nil {
		log.Fatal(err)
	}

	compileMaster(courseDir)

	for _, chapter := range chapters {
		if filepath.Ext(chapter.Name()) == ".tex" {
			compileChapter(chapter.Name(), courseDir)
		}
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
