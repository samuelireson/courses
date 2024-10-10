#!/usr/bin/bash

if [ ! -d "$1" ]; then
	echo "$1 is not a valid course";
	exit 1
fi

for chapter in $1/chapters/*; do
	./scripts/convert-chapter.pl $chapter
done
