#!/bin/bash

cd ..

output="README.md"
[[ -f $output ]] && mv $output $output.bak

cat "scripts/firstpart.md" >$output

for folder in "."/*; do
	if [ -d "$folder" ]; then
		if [ -f "$folder/README.md" ]; then
			sed '/^#/ s/^/##/' "$folder/README.md" >>"$output"
			echo "" >>"$output"
		fi
	fi
done

cat "scripts/lastpart.md" >>$output

md-table-of-contents "$output" >temp_table_of_contents.md
awk 'FNR==NR { a[i++] = $0; next } /!tableofcontents/ { for (j=0; j<i; j++) print a[j]; next } 1' temp_table_of_contents.md "$output" >temp_readme.md
mv temp_readme.md "$output"
rm temp_table_of_contents.md
