/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package cmd

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strings"
)

type bibliographyItem struct {
	author string
	title  string
}

type bibliography map[string]bibliographyItem

var keyMatch = regexp.MustCompile(`{[a-z]*,`)
var valueMatch = regexp.MustCompile(`{.*}`)

func cleanMatch(rawMatch string) string {
	cleanMatch := strings.TrimPrefix(rawMatch, "{")
	cleanMatch = strings.TrimSuffix(cleanMatch, "}")
	cleanMatch = strings.TrimSuffix(cleanMatch, ",")

	return cleanMatch
}

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

		switch {
		case strings.Contains(line, "@"):
			currentKey = cleanMatch(keyMatch.FindString(line))
		case strings.Contains(line, "author"):
			currentAuthor = cleanMatch(valueMatch.FindString(line))
		case strings.Contains(line, "title"):
			currentTitle = cleanMatch(valueMatch.FindString(line))
		}

		bib[currentKey] = bibliographyItem{
			author: currentAuthor,
			title:  currentTitle,
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return bib
}

var citationMatch = regexp.MustCompile(`\\cite\{(.*?)\}`)

func convertCitationsToFootnotes(bib bibliography, content string) string {

	citations := citationMatch.FindAllString(content, -1)
	var citationContent []string
	var citationContentItem string

	for _, citation := range citations {
		citation = strings.TrimPrefix(citation, "\\cite{")
		citation = strings.TrimSuffix(citation, "}")
		citationBibItem := bib[citation]

		citationContentItem = "[^" + citation + "]: <em> " + citationBibItem.title + " </em> -- " + citationBibItem.author
		citationContent = append(citationContent, citationContentItem)
		slices.Sort(citationContent)
		citationContent = slices.Compact(citationContent)
	}

	content = citationMatch.ReplaceAllString(content, string("[^$1]"))
	content += strings.Join(citationContent, "\n")

	return content
}
