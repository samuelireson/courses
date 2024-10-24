/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package convert

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type regexPattern struct {
	captureGroup *regexp.Regexp
	replacement  string
}

type stringPattern struct {
	old string
	new string
}

var header = `---
title: $1
---
import Aside from '@components/Aside/Aside.tsx';
import Comments from '@components/Comments/Comments.tsx';
import { Tabs, TabItem, LinkButton } from '@astrojs/starlight/components';

`

var basicRegexPatterns = []regexPattern{
	// document organisation
	{regexp.MustCompile(`\\chapter\{(.*?)\}`), header},
	{regexp.MustCompile(`\\section\{(.*?)\}`), "## $1"},
	{regexp.MustCompile(`\\subsection\{(.*?)\}`), "### $1"},

	// theorem environments
	{regexp.MustCompile(`\\end\{(definition|theorem|lemma|proposition|corollary|example|nonexample|notation|remark)\}`), "</Aside>"},

	// fonts and ligatures
	{regexp.MustCompile(`\\textbf\{(.*?)\}`), "<b> $1 </b>"},
	{regexp.MustCompile(`\\textit\{(.*?)\}`), "<em> $1 </em>"},
}

var stringPatterns = []stringPattern{
	//document organisation
	{"\\begin{chout}", "<div style='text-align: center'><em>"},
	{"\\end{chout}", "</em></div>"},

	// maths environments
	{"\\begin{align*}", "$$\n\\begin{align*}"},
	{"\\end{align*}", "\\end{align*}\n$$"},

	// exercises
	{"\\begin{exercise}", "<Tabs>"},
	{"\\end{exercise}", "</Tabs>"},
	{"\\begin{problem}", "<TabItem label='Problem'>"},
	{"\\begin{solution}", "<TabItem label='Solution'>"},
	{"\\end{problem}", "</TabItem>"},
	{"\\end{solution}", "</TabItem>"},

	// theorem environments
	{"\\begin{corollary}", "<Aside type='result' title='Corollary' >"},
	{"\\begin{definition}", "<Aside type='definition' title='Definition' >"},
	{"\\begin{example}", "<Aside type='example' title='Example' >"},
	{"\\begin{lemma}", "<Aside type='result' title='Lemma' >"},
	{"\\begin{nonexample}", "<Aside type='example' title='Nonexample' >"},
	{"\\begin{notation}", "<Aside type='comment' title='Notation' >"},
	{"\\begin{proposition}", "<Aside type='result' title='Proposition' >"},
	{"\\begin{remark}", "<Aside type='comment' title='Remark' >"},
	{"\\begin{theorem}", "<Aside type='result' title='Theorem' >"},

	// badges
	{"\\basic", ":badge[Basic]{variant=success}"},
	{"\\intermediate", ":badge[Intermediate]{variant=warning}"},
	{"\\challenging", ":badge[Challenging]{variant=danger}"},

	// unordered lists
	{"\\begin{itemize}", ""},
	{"\\item", "-"},
	{"\\end{itemize}", ""},

	// fonts and ligatures
	{"`", "'"},
}

func convertTeXToMDX(content string) string {
	for _, element := range stringPatterns {
		content = strings.ReplaceAll(content, element.old, element.new)
	}

	for _, element := range basicRegexPatterns {
		content = element.captureGroup.ReplaceAllString(content, element.replacement)
	}
	return content
}

func addDownloadLinks(content, inputPath string) string {

	inputPathSections := strings.Split(inputPath, string(filepath.Separator))

	courseName := "/" + filepath.Join(inputPathSections[0], inputPathSections[1])
	chapterName := strings.TrimSuffix(inputPathSections[3], filepath.Ext(inputPathSections[3])) + ".pdf"
	chapterDownloadPath := filepath.Join("courses", courseName, chapterName)
	courseDownloadPath := filepath.Join("courses", courseName, "master.pdf")

	downloadLinkTemplate := fmt.Sprintf("<div style='display: flex; justify-content: space-around;'>\n\t<LinkButton target=\"_blank\" href=\"%s\" variant=\"secondary\" icon=\"document\" >Download</LinkButton>\n\t<LinkButton target=\"_blank\" href=\"%s\" variant=\"primary\" icon=\"open-book\" >Download</LinkButton>\n</div>", chapterDownloadPath, courseDownloadPath)

	content = strings.Replace(content, "\n\n", "\n\n"+downloadLinkTemplate, 1)

	return content
}

func addComments(content string) string {
	content += "\n<Comments client:only=\"react\" />"
	return content
}
