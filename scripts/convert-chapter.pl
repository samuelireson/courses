#!/usr/bin/perl
use strict;
use warnings;

# Check if correct number of arguments are passed
if (@ARGV != 1) {
	die "Usage: $0 input.tex\n";
}

my ($input_file) = @ARGV;

my $output_file = $input_file;
$output_file =~ s/\.tex/.mdx/;

my $course = $input_file;
$course =~ s/notes\/(\w+?)\/.*$/$1/;

my $chapter = $input_file;
$chapter =~ s/notes\/\w*?\/chapters\/(.*?).tex/\u$1/;
$chapter =~ s/-/ /g;

my $inside_table = 0;
my $cols;

# Open input and output files
open(my $in, '<', $input_file) or die "Cannot open $input_file: $!";
open(my $out, '>', $output_file) or die "Cannot open $output_file: $!";

print $out "---
title: $chapter
---

import { Aside } from '\@components';
import { Tabs, TabItem } from '\@astrojs/starlight/components';

";

while (my $line = <$in>) {
	# Remove \n from EOL
	chomp $line;

	# Convert chout env
	$line =~ s/\\begin\{chout\}/<div style="text-align: center"><em>/;
	$line =~ s/\\end\{chout\}/<\/em><\/div>/;

	# Convert LaTeX section headings to Markdown headers
	next if $line =~ /\\chapter/;
	$line =~ s/\\section\{(.+?)\}/## $1/;
	$line =~ s/\\subsection\{(.+?)\}/### $1/;
	$line =~ s/\\basic/:badge[Basic]{variant=success}/;
	$line =~ s/\\intermediate/:badge[Intermediate]{variant=warning}/;
	$line =~ s/\\challenging/:badge[Challenging]{variant=danger}/;

	# Convert theorem like environments
	$line =~ s/\\begin\{definition\}/<Aside type='definition' title='Definition' >/;
	$line =~ s/\\begin\{(theorem|lemma|proposition|corollary)\}/<Aside type='result' title='\u$1' >/;
	$line =~ s/\\begin\{(example|nonexample)\}/<Aside type='example' title='\u$1' >/;
	$line =~ s/\\begin\{(notation|remark)\}/<Aside type='comment' title='\u$1' >/;
	$line =~ s/\\end\{(definition|theorem|lemma|proposition|corollary|example|nonexample|notation|remark)\}/<\/Aside>/;

	# Convert exercises
	$line =~ s/\\begin\{exercise\}/<Tabs>/;
	$line =~ s/\\end\{exercise\}/<\/Tabs>/;
	$line =~ s/\\begin\{problem\}/<TabItem label="Problem">/;
	$line =~ s/\\begin\{solution\}/<TabItem label="Solution">/;
	$line =~ s/\\end\{(problem|solution)\}/<\/TabItem>/;

	# Convert math environments
	$line =~ s/(\\begin\{align\*\})/\$\$\n$1/;
	$line =~ s/(\\end\{align\*\})/$1\n\$\$/;

	# Change accent text
	$line =~ s/\\textbf\{(.+?)\}/**$1**/g;
	$line =~ s/\\textit\{(.+?)\}/*$1*/g;

	# Convert tables
	$line =~ s/\\begin\{center\}/<center>/;
	$line =~ s/\\end\{center\}/<\/center>/;
	if ($line =~ /\\begin\{tabular\}\{(?<cols>.*?)\}/) {
		$cols = $+{cols};
		$cols =~ s/(c|l|r)/---/g;
		$cols =~ s/^/\n/;
		$inside_table = 1;
		next;
	}
	if ($inside_table) {
		$line =~ s/\&/|/g;
		$line =~ s/\\\\/$cols/;
		next if $line =~ /\\hline/;
	}
	if ($line =~ /\\end\{tabular\}/) {
		$inside_table = 0;
		next;
	}


	# Convert LaTeX quotations
	$line =~ s/`(.*?)'/'$1'/g;

	# Remove LaTeX comment lines (lines starting with %)
	next if $line =~ /^\s*%/;

	# Remove citations
	$line =~ s/\\cite\{.+?\}//g;
	$line =~ s/\\ref\{.+?\}//g;
	$line =~ s/\\label\{.+?\}//g;

	# Custom macros redefinition
	$line =~ s/\\defined\{(.*?)\}/**$1**/gs;

	print $out "$line\n";
	$cols = '';
}

# Close the files
close($in);
close($out);

print "Conversion complete!\n";

system "mv $output_file ./site/src/content/docs/$course/";

print "File moved!\n";
