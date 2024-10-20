/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type bibliographyItem struct {
	author string
	title  string
}

type bibliography map[string]bibliographyItem

var keyMatch = regexp.MustCompile(`{[a-z]*,`)
var valueMatch = regexp.MustCompile(`{.*}`)

func parseBibliography(bibPath string) bibliography {
	var bib bibliography
	bib = make(bibliography)

	var currentKey string
	var currentAuthor string
	var currentTitle string

	fi, err := os.Open(bibPath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "@") {
			currentKey = keyMatch.FindString(line)
			currentKey = strings.TrimPrefix(currentKey, "{")
			currentKey = strings.TrimSuffix(currentKey, ",")
		} else if strings.Contains(line, "author") {
			currentAuthor = valueMatch.FindString(line)
			currentAuthor = strings.TrimPrefix(currentAuthor, "{")
			currentAuthor = strings.TrimSuffix(currentAuthor, "}")
		} else if strings.Contains(line, "title") {
			currentTitle = valueMatch.FindString(line)
			currentTitle = strings.TrimPrefix(currentTitle, "{")
			currentTitle = strings.TrimSuffix(currentTitle, "}")
		}

		bib[currentKey] = bibliographyItem{
			author: currentAuthor,
			title:  currentTitle,
		}

	}

	for key, item := range bib {
		fmt.Printf("Key: %s -> Author: %s, Title: %s\n", key, item.author, item.title)
	}

	return bib
}

func convertCitationsToFootnotes(bib bibliography) {

}

// @book{donaldson,
//   author        = {Donaldson, Simon},
//   publisher     = {Oxford University Press},
//   series        = {Oxford Graduate Texts in Mathematics},
//   title         = {Riemann Surfaces},
//   volume        = {22},
//   year          = {2011}
// }
