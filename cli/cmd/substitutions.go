/*
Copyright © 2024 Samuel Ireson samuelireson@gmail.com
*/
package cmd

import (
	"regexp"
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
import { Aside } from '\@components';
import { Tabs, TabItem } from '\@astrojs/starlight/components';
`

var basicRegexPatterns = []regexPattern{
	// document organisation
	{regexp.MustCompile(`\\chapter\{(.*?)\}`), header},
	{regexp.MustCompile(`\\section\{(.*?)\}`), "## $1"},
	{regexp.MustCompile(`\\subsection\{(.*?)\}`), "### $1"},

	// theorem environments
	{regexp.MustCompile(`\\begin\{definition\}`), "<Aside type='definition' title='definition' >"},
	{regexp.MustCompile(`\\begin\{(theorem|lemma|proposition|corollary)\}`), "<Aside type='result' title='$1' >"},
	{regexp.MustCompile(`\\begin\{(example|nonexample)\}`), "<Aside type='example' title='$1' >"},
	{regexp.MustCompile(`\\begin\{(notation|remark)\}`), "<Aside type='comment' title='$1' >"},
	{regexp.MustCompile(`\\end\{(definition|theorem|lemma|proposition|corollary|example|nonexample|notation|remark)\}`), "</Aside>"},

	// badges
	{regexp.MustCompile(`\\basic`), ":badge[Basic]{variant=success}"},
	{regexp.MustCompile(`\\intermediate`), ":badge[Intermediate]{variant=warning}"},
	{regexp.MustCompile(`\\challenging`), ":badge[Challenging]{variant=danger}"},

	// fonts and ligatures
	{regexp.MustCompile(`\\textbf\{(.*?)\}`), "* $1 *"},
	{regexp.MustCompile(`\\textit\{(.*?)\}`), "** $1 **"},
	{regexp.MustCompile("`"), "'"},
}

var stringPatterns = []stringPattern{
	//document organisation
	{"\\begin{chout}", "<div style='text-align: center'><em>"},
	{"\\end{chout}", "</em></div>"},

	// maths environments
	{"\\begin{align*})", "$$$$\n\\begin{align*}"},
	{"\\end{align*}", "\\end{align*}\n$$$$"},

	// exercises
	{"\\begin{exercise}", "<Tabs>"},
	{"\\end{exercise}", "</Tabs>"},
	{"\\begin{problem}", "<TabItem label='Problem'>"},
	{"\\begin{solution}", "<TabItem label='Solution'>"},
	{"\\end{problem}", "</TabItem>"},
	{"\\end{solution}", "</TabItem>"},
}
