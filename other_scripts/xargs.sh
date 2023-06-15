#!/bin/bash

cat example.txt | xargs
cat example.txt | xargs -n 3
echo 'abcAdefAghiAjklAmnoApqr' | xargs -d A
echo 'abcAdefAghiAjklAmnoApqr' | xargs -d A -n 2
find . -name '*.sh' -print0 | xargs -0 grep -L images
cat args.txt | xargs -I {} ./cecho.sh -p {} -l
cat args.txt | xargs -I arg ./cecho.sh -p arg -l
find . -name '*.txt' | xargs wc -l
find . -name '*.txt' -print0 | xargs -0 wc -l
